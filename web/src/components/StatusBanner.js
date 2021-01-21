import { useContext } from "react";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import Typography from "@material-ui/core/Typography";

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

  return (
    <div style={{ marginTop: "5%", padding: "5%" }}>
      <center>
        <Typography variant="h3">HashMap is online</Typography>
        <Typography variant="h5">
          {getNumberOfConnectedDevice(appStore.devices)} devices connected
        </Typography>
      </center>
    </div>
  );
});

export default StatusBanner;
