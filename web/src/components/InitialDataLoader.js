import { useContext, useEffect, useState } from "react";
import axios from "axios";
import Typography from "@material-ui/core/Typography";
import LinearProgress from "@material-ui/core/LinearProgress";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";

const InitialDataLoader = observer(({}) => {
  const appStore = useContext(AppStoreContext);

  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    setIsLoading(true);
    axios({
      method: "GET",
      url: "http://localhost:4209/ohw/api/devices",
    })
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
