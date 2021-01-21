import { useContext } from "react";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import Typography from "@material-ui/core/Typography";

const StatusBanner = observer(({}) => {
  const appStore = useContext(AppStoreContext);

  return (
    <div style={{ marginTop: "5%", padding: "5%" }}>
      <center>
        <Typography variant="h3">Network is online</Typography>
        <Typography variant="h5">
          {appStore.devices.length} devices connected
        </Typography>
      </center>
    </div>
  );
});

export default StatusBanner;
