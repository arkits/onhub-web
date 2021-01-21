import { useContext } from "react";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import Typography from "@material-ui/core/Typography";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import Grid from "@material-ui/core/Grid";
import WifiIcon from "@material-ui/icons/Wifi";
import SettingsInputHdmiIcon from "@material-ui/icons/SettingsInputHdmi";
import PermScanWifiIcon from "@material-ui/icons/PermScanWifi";

function DeviceCard({ device }) {
  const renderConnectionIcon = (connectionType) => {
    if (connectionType === "WIRELESS") {
      return <WifiIcon fontSize="large" />;
    } else if (connectionType === "WIRED") {
      return <SettingsInputHdmiIcon fontSize="large" />;
    } else {
      return <PermScanWifiIcon fontSize="large" />;
    }
  };

  return (
    <div style={{ paddingBottom: "10px" }}>
      <Card>
        <CardContent>
          <Grid container spacing="3">
            <Grid item xs={2}>
              <center>
                {renderConnectionIcon(device?.connectionType)}
                <Typography variant="subtitle1" color="textPrimary">
                  {device?.connectionType}
                </Typography>
              </center>
            </Grid>
            <Grid item xs={10}>
              <Typography variant="subtitle1" color="textPrimary">
                {device?.friendlyName}
              </Typography>
              <Typography variant="body2" color="textSecondary">
                {device?.id}
              </Typography>
              <pre>{JSON.stringify(device, null, 4)}</pre>
            </Grid>
          </Grid>
        </CardContent>
      </Card>
    </div>
  );
}

function DeviceCards({ devices }) {
  let elements = [];

  devices.forEach((device, index) => {
    elements.push(<DeviceCard key={index} device={device} />);
  });

  return elements;
}

const DeviceCardListing = observer(({}) => {
  const appStore = useContext(AppStoreContext);
  return (
    <div>
      <div style={{ paddingBottom: "20px" }}>
        <Typography variant="h6">Connected Devices</Typography>
      </div>
      <DeviceCards devices={appStore.devices} />
    </div>
  );
});

export default DeviceCardListing;
