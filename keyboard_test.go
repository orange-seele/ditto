package main

import (
	"testing"

	evdev "github.com/gvalkov/golang-evdev"
)

func TestIsKeyboardDevice_happyPath(t *testing.T) {
	dev := &evdev.InputDevice{
		Capabilities: map[evdev.CapabilityType][]evdev.CapabilityCode{
			{Type: evdev.EV_KEY}: {
				{Code: evdev.KEY_A},
			},
		},
	}
	if !isKeyboardDevice(dev) {
		t.Error("expected device with KEY_A to be a keyboard")
	}
}

func TestIsKeyboardDevice_missingKeyA(t *testing.T) {
	dev := &evdev.InputDevice{
		Capabilities: map[evdev.CapabilityType][]evdev.CapabilityCode{
			{Type: evdev.EV_KEY}: {
				{Code: evdev.KEY_ESC},
				{Code: evdev.KEY_ENTER},
			},
		},
	}
	if isKeyboardDevice(dev) {
		t.Error("expected device without KEY_A to NOT be a keyboard")
	}
}

func TestIsKeyboardDevice_noEvKey(t *testing.T) {
	dev := &evdev.InputDevice{
		Capabilities: map[evdev.CapabilityType][]evdev.CapabilityCode{
			{Type: evdev.EV_ABS}: {
				{Code: 0}, // ABS_X
				{Code: 1}, // ABS_Y
			},
		},
	}
	if isKeyboardDevice(dev) {
		t.Error("expected device with only EV_ABS to NOT be a keyboard")
	}
}

func TestIsKeyboardDevice_emptyCapabilities(t *testing.T) {
	dev := &evdev.InputDevice{
		Capabilities: map[evdev.CapabilityType][]evdev.CapabilityCode{},
	}
	if isKeyboardDevice(dev) {
		t.Error("expected device with empty capabilities to NOT be a keyboard")
	}
}

func TestIsKeyboardDevice_mixedCapabilities(t *testing.T) {
	dev := &evdev.InputDevice{
		Capabilities: map[evdev.CapabilityType][]evdev.CapabilityCode{
			{Type: evdev.EV_MSC}: {
				{Code: 4},
			},
			{Type: evdev.EV_KEY}: {
				{Code: evdev.KEY_A},
				{Code: evdev.KEY_B},
			},
		},
	}
	if !isKeyboardDevice(dev) {
		t.Error("expected device with EV_KEY + KEY_A among other cap types to be a keyboard")
	}
}
