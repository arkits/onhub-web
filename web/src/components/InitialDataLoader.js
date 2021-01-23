import { useContext, useEffect, useState } from "react";
import { Typography, LinearProgress } from "@material-ui/core";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import { getDevices } from "../api/DevicesApi";

const InitialDataLoader = observer(({}) => {
  const appStore = useContext(AppStoreContext);

  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    setIsLoading(true);

    // Get Initial Devices
    getDevices()
      .then(function (response) {
        appStore.devices = response?.data;
        appStore.isInitialLoadComplete = true;
        setIsLoading(false);
      })
      .catch(function (err) {
        console.error(err);
        setIsLoading(false);
      });
  }, []);

  if (isLoading) {
    return (
      <div style={{ marginTop: "10%" }}>
        <center>
          <Typography variant="h5">Loading Initial Data...</Typography> <br />
          <LinearProgress color="secondary" />
        </center>
      </div>
    );
  } else {
    return null;
  }
});

export default InitialDataLoader;
