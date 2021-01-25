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

  const renderNetworkMetrics = (networkMetrics) => {
    if (networkMetrics) {
      return `${formatBytes(
        networkMetrics[0]?.network_metrics?.groupTraffic?.transmitSpeedBps
      )} upload - ${formatBytes(
        networkMetrics[0]?.network_metrics?.groupTraffic?.receiveSpeedBps
      )} download`;
    } else {
      return null;
    }
  };

  return (
    <div style={{ marginTop: "5%", padding: "5%" }}>
      <center>
        <Typography variant="h3">HashMap is online</Typography>
        <Typography variant="h5">
          {getNumberOfConnectedDevice(appStore.devices)} devices connected{" "}
          <br></br>
          {renderNetworkMetrics(appStore.networkMetrics.data)}
        </Typography>
      </center>
    </div>
  );
});

export default StatusBanner;
