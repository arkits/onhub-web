import { useContext } from "react";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import Typography from "@material-ui/core/Typography";
import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import Button from "@material-ui/core/Button";

function DeviceCard({ device }) {
  return (
    <div style={{ paddingBottom: "10px" }}>
      <Card>
        <CardContent>
          <Typography color="textSecondary">{device?.id}</Typography>
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
        <center>
          <Typography variant="overline">Connected Devices</Typography>
        </center>
      </div>
      <DeviceCards devices={appStore.devices} />
    </div>
  );
});

export default DeviceCardListing;
