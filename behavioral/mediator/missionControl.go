package main

type missionControl struct {
	launchPadAvailable bool
	rocketQueue        []irocket
}

func newMissionControl() *missionControl {
	return &missionControl{
		launchPadAvailable: true,
	}
}

func (mc *missionControl) canMoveToLaunchPad(r irocket) bool {
	if mc.launchPadAvailable {
		mc.launchPadAvailable = false
		return true
	}
	mc.rocketQueue = append(mc.rocketQueue, r)
	return false
}

func (mc *missionControl) notifySuccessfulLaunch() {
	if !mc.launchPadAvailable {
		mc.launchPadAvailable = true
	}
	if len(mc.rocketQueue) > 0 {
		nextRocket := mc.rocketQueue[0]
		mc.rocketQueue = mc.rocketQueue[1:]
		nextRocket.moveToLaunchPad()
	}
}
