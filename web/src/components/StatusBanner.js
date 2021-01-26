import { useContext } from "react";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import Typography from "@material-ui/core/Typography";
import { formatBytes } from "../utils/utils";

const StatusBanner = observer(({}) => {
  const appStore = useContext(AppStoreContext);

  const getNumberOfConnectedDevice = (devices) => {
    let numberOfConnectedDevices = 0;
    devices.forEach((device) => {
      if (device?.ipAddresses !== null) {
        numberOfConnectedDevices++;
      }
    });
    return numberOfConnectedDevices;
  };

  const renderNetworkMetric = (networkMetrics) => {
    if (networkMetrics) {
      let prettyUpload = formatBytes(
        networkMetrics[0]?.network_metrics?.groupTraffic?.transmitSpeedBps
      );

      let prettyDownload = formatBytes(
        networkMetrics[0]?.network_metrics?.groupTraffic?.receiveSpeedBps
      );

      return (
        <Typography
          variant="h5"
          color="textSecondary"
          style={{ fontWeight: 700 }}
        >
          {prettyDownload} download <br />
          {prettyUpload} upload<br />
        </Typography>
      );
    } else {
      return null;
    }
  };

  return (
    <div>
      <Typography variant="h3" style={{ fontWeight: 700 }}>
        {getNumberOfConnectedDevice(appStore?.devices)} devices connected{" "}
      </Typography>
      {renderNetworkMetric(appStore?.networkMetrics?.data)}
    </div>
  );
});

export default StatusBanner;
