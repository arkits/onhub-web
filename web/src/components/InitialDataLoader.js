import { useContext, useEffect, useState } from "react";
import {
  Typography,
  LinearProgress,
  Card,
  CardContent,
} from "@material-ui/core";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import { getDevices } from "../api/DevicesApi";

function ErrorBanner({ apiError }) {
  if (apiError != null) {
    return (
      <>
        <Card style={{ margin: 20, background: "red" }}>
          <CardContent>
            <Typography variant="h5">
              Server Error: {apiError.message}
            </Typography>
            <Typography variant="body1">
              An Error with the Server has occured. Here are some details -
            </Typography>
            <pre>{JSON.stringify(apiError, null, 4)}</pre>
          </CardContent>
        </Card>
      </>
    );
  } else {
    return null;
  }
}

const InitialDataLoader = observer(({}) => {
  const appStore = useContext(AppStoreContext);

  const [isLoading, setIsLoading] = useState(false);

  const [apiError, setApiError] = useState(null);

  useEffect(() => {
    setIsLoading(true);
    getDevices()
      .then(function (response) {
        appStore.devices = response?.data;
        appStore.isInitialLoadComplete = true;
        setIsLoading(false);
      })
      .catch(function (err) {
        console.error(err);
        setIsLoading(false);
        setApiError(err);
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
    return <ErrorBanner apiError={apiError} />;
  }
});

export default InitialDataLoader;
