package spine

import (
	"errors"
	"fmt"
	"reflect"
	"sync/atomic"

	"github.com/DerAndereAndi/eebus-go/spine/model"
	"github.com/ahmetb/go-linq/v3"
)

type SubscriptionManager interface {
	AddSubscription(localDevice *DeviceLocalImpl, remoteDevice *DeviceRemoteImpl, data model.SubscriptionManagementRequestCallType) error
	RemoveSubscription(data model.SubscriptionManagementDeleteCallType, remoteDevice *DeviceRemoteImpl) error
	Subscriptions(remoteDevice *DeviceRemoteImpl) []*SubscriptionEntry
	SubscriptionsOnFeature(featureAddress model.FeatureAddressType) []*SubscriptionEntry
}

type SubscriptionEntry struct {
	id            uint64
	serverFeature FeatureLocal
	clientFeature *FeatureRemoteImpl
}

type SubscriptionManagerImpl struct {
	subscriptionNum     uint64
	subscriptionEntries []*SubscriptionEntry
	// TODO: add persistence
}

func NewSubscriptionManager() SubscriptionManager {
	c := &SubscriptionManagerImpl{
		subscriptionNum: 0,
	}

	return c
}

// is sent from the client (remote device) to the server (local device)
func (c *SubscriptionManagerImpl) AddSubscription(localDevice *DeviceLocalImpl, remoteDevice *DeviceRemoteImpl, data model.SubscriptionManagementRequestCallType) error {

	serverFeature := localDevice.FeatureByAddress(data.ServerAddress)
	if serverFeature == nil {
		return fmt.Errorf("server feature '%s' in local device '%s' not found", data.ServerAddress, *localDevice.Address())
	}
	if err := c.checkRole(serverFeature, model.RoleTypeServer, *data.ServerFeatureType); err != nil {
		return err
	}

	clientFeature := remoteDevice.FeatureByAddress(data.ClientAddress)
	if clientFeature == nil {
		return fmt.Errorf("client feature '%s' in remote device '%s' not found", data.ClientAddress, *remoteDevice.Address())
	}
	if err := c.checkRole(clientFeature, model.RoleTypeClient, *data.ServerFeatureType); err != nil {
		return err
	}

	subscriptionEntry := &SubscriptionEntry{
		id:            c.subscriptionId(),
		serverFeature: serverFeature,
		clientFeature: clientFeature,
	}

	// TOV-TODO: check if subscription already exists
	c.subscriptionEntries = append(c.subscriptionEntries, subscriptionEntry)

	payload := EventPayload{
		Ski:        remoteDevice.ski,
		EventType:  EventTypeSubscriptionChange,
		ChangeType: ElementChangeAdd,
		Data:       data,
		Feature:    clientFeature,
	}
	Events.Publish(payload)

	// TOV-TODO: Send heartbeat to the feature which subscribed to DeviceDiagnostic

	return nil
}

func (c *SubscriptionManagerImpl) RemoveSubscription(data model.SubscriptionManagementDeleteCallType, remoteDevice *DeviceRemoteImpl) error {
	// TODO: test this!!!

	var newSubscriptionEntries []*SubscriptionEntry

	// according to the spec 7.4.4
	// a. The absence of "subscriptionDelete. clientAddress. device" SHALL be treated as if it was
	//    present and set to the sender's "device" address part.
	// b. The absence of "subscriptionDelete. serverAddress. device" SHALL be treated as if it was
	//    present and set to the recipient's "device" address part.

	clientAddress := data.ClientAddress
	if data.ClientAddress.Device == nil {
		clientAddress.Device = remoteDevice.Address()
	}

	for _, item := range c.subscriptionEntries {
		if !reflect.DeepEqual(item.clientFeature.Address(), clientAddress) {
			newSubscriptionEntries = append(newSubscriptionEntries, item)
		}
	}

	if len(newSubscriptionEntries) == len(c.subscriptionEntries) {
		return errors.New("could not find requested SubscriptionId to be removed")
	}

	c.subscriptionEntries = newSubscriptionEntries

	// TOV-TODO: stop heartbeat for remote device when it has no subscription to DeviceDiagnostic anymore
	return nil
}

func (c *SubscriptionManagerImpl) Subscriptions(remoteDevice *DeviceRemoteImpl) []*SubscriptionEntry {
	var result []*SubscriptionEntry

	linq.From(c.subscriptionEntries).WhereT(func(s *SubscriptionEntry) bool {
		return s.clientFeature.Device().Ski() == remoteDevice.Ski()
	}).ToSlice(&result)

	return result
}

func (c *SubscriptionManagerImpl) SubscriptionsOnFeature(featureAddress model.FeatureAddressType) []*SubscriptionEntry {
	var result []*SubscriptionEntry

	linq.From(c.subscriptionEntries).WhereT(func(s *SubscriptionEntry) bool {
		return reflect.DeepEqual(*s.serverFeature.Address(), featureAddress)
	}).ToSlice(&result)

	return result
}

func (c *SubscriptionManagerImpl) subscriptionId() uint64 {
	i := atomic.AddUint64(&c.subscriptionNum, 1)
	return i
}

func (c *SubscriptionManagerImpl) checkRole(feature Feature, role model.RoleType, featureType model.FeatureTypeType) error {
	if feature.Role() != model.RoleTypeSpecial && feature.Role() != role {
		return fmt.Errorf("found feature '%s' is not matching required role '%s'", feature.Type(), role)
	}

	if feature.Type() != featureType {
		return fmt.Errorf("found feature '%s' is not matching required type '%s'", feature.Type(), featureType)
	}

	return nil
}
