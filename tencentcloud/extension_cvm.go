package tencentcloud

const (
	CVM_CHARGE_TYPE_PREPAID  = "PREPAID"
	CVM_CHARGE_TYPE_POSTPAID = "POSTPAID_BY_HOUR"
	CVM_CHARGE_TYPE_SPOTPAID = "SPOTPAID"

	CVM_INTERNET_CHARGE_TYPE_BANDWIDTH_PREPAID  = "BANDWIDTH_PREPAID"
	CVM_INTERNET_CHARGE_TYPE_BANDWIDTH_POSTPAID = "BANDWIDTH_POSTPAID_BY_HOUR"
	CVM_INTERNET_CHARGE_TYPE_BANDWIDTH_PACKAGE  = "BANDWIDTH_PACKAGE"
	CVM_INTERNET_CHARGE_TYPE_TRAFFIC_POSTPAID   = "TRAFFIC_POSTPAID_BY_HOUR"

	CVM_STATUS_RUNNING       = "RUNNING"
	CVM_STATUS_STOPPED       = "STOPPED"
	CVM_STATUS_SHUTDOWN      = "SHUTDOWN"
	CVM_STATUS_TERMINATING   = "TERMINATING"
	CVM_STATUS_LAUNCH_FAILED = "LAUNCH_FAILED"

	CVM_LATEST_OPERATION_STATE_OPERATING = "OPERATING"
	CVM_LATEST_OPERATION_STATE_SUCCESS   = "SUCCESS"
	CVM_LATEST_OPERATION_STATE_FAILED    = "FAILED"

	CVM_PREPAID_RENEW_FLAG_NOTIFY_NOTIFY_AND_AUTO_RENEW    = "NOTIFY_AND_AUTO_RENEW"
	CVM_PREPAID_RENEW_FLAG_NOTIFY_AND_MANUAL_RENEW         = "NOTIFY_AND_MANUAL_RENEW"
	CVM_PREPAID_RENEW_FLAG_DISABLE_NOTIFY_AND_MANUAL_RENEW = "DISABLE_NOTIFY_AND_MANUAL_RENEW"

	CVM_DISK_TYPE_LOCAL_BASIC   = "LOCAL_BASIC"
	CVM_DISK_TYPE_LOCAL_SSD     = "LOCAL_SSD"
	CVM_DISK_TYPE_CLOUD_BASIC   = "CLOUD_BASIC"
	CVM_DISK_TYPE_CLOUD_SSD     = "CLOUD_SSD"
	CVM_DISK_TYPE_CLOUD_PREMIUM = "CLOUD_PREMIUM"

	CVM_PLACEMENT_GROUP_TYPE_HOST = "HOST"
	CVM_PLACEMENT_GROUP_TYPE_SW   = "SW"
	CVM_PLACEMENT_GROUP_TYPE_RACK = "RACK"

	ZONE_STATE_AVAILABLE   = "AVAILABLE"
	ZONE_STATE_UNAVAILABLE = "UNAVAILABLE"

	CVM_NOT_FOUND_ERROR = "InvalidInstanceId.NotFound"

	CVM_SPOT_INSTANCE_TYPE_ONE_TIME = "ONE-TIME"

	CVM_MARKET_TYPE_SPOT = "spot"
)

var CVM_CHARGE_TYPE = []string{
	CVM_CHARGE_TYPE_PREPAID,
	CVM_CHARGE_TYPE_POSTPAID,
	CVM_CHARGE_TYPE_SPOTPAID,
}

var CVM_INTERNET_CHARGE_TYPE = []string{
	CVM_INTERNET_CHARGE_TYPE_BANDWIDTH_PREPAID,
	CVM_INTERNET_CHARGE_TYPE_BANDWIDTH_POSTPAID,
	CVM_INTERNET_CHARGE_TYPE_BANDWIDTH_PACKAGE,
	CVM_INTERNET_CHARGE_TYPE_TRAFFIC_POSTPAID,
}

var CVM_PREPAID_PERIOD = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36}

var CVM_PREPAID_RENEW_FLAG = []string{
	CVM_PREPAID_RENEW_FLAG_NOTIFY_NOTIFY_AND_AUTO_RENEW,
	CVM_PREPAID_RENEW_FLAG_NOTIFY_AND_MANUAL_RENEW,
	CVM_PREPAID_RENEW_FLAG_DISABLE_NOTIFY_AND_MANUAL_RENEW,
}

var CVM_DISK_TYPE = []string{
	CVM_DISK_TYPE_LOCAL_BASIC,
	CVM_DISK_TYPE_LOCAL_SSD,
	CVM_DISK_TYPE_CLOUD_BASIC,
	CVM_DISK_TYPE_CLOUD_SSD,
	CVM_DISK_TYPE_CLOUD_PREMIUM,
}

var CVM_PLACEMENT_GROUP_TYPE = []string{
	CVM_PLACEMENT_GROUP_TYPE_HOST,
	CVM_PLACEMENT_GROUP_TYPE_SW,
	CVM_PLACEMENT_GROUP_TYPE_RACK,
}

var CVM_SPOT_INSTANCE_TYPE = []string{
	CVM_SPOT_INSTANCE_TYPE_ONE_TIME,
}
