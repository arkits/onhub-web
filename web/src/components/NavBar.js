import { AppBar, Button, Toolbar, Typography } from "@material-ui/core/";
import { makeStyles } from "@material-ui/core/styles";
import { useState, useEffect } from "react";
import { getVersion } from "../api/DebugApi";
import NetworkMetricsStatusDialog from "./NetworkMetricsStatusDialog";

const useStyles = makeStyles((theme) => ({
  titleText: {
    display: "flex",
    flexGrow: "1",
    fontFamily: "IBM Plex Mono",
    fontWeight: "bold",
    fontStyle: "italic",
  },
  subtitleText: {
    fontFamily: "IBM Plex Mono",
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
  },
  buttonStyle: {
    fontFamily: "Inter",
    fontWeight: "bold",
    textTransform: "none",
  },
}));

function NavBar() {
  const classes = useStyles();

  const [appVersion, setAppVersion] = useState("0.0.1");

  useEffect(() => {
    getVersion()
      .then(function (response) {
        setAppVersion(response?.data?.version);
      })
      .catch(function (err) {
        console.error(err);
      });
  }, []);

  return (
    <div>
      <AppBar position="fixed" className={classes.appBar} color="primary">
        <Toolbar>
          <Typography className={classes.titleText} variant="h5" noWrap>
            ./onhub-web
          </Typography>
          <NetworkMetricsStatusDialog />
          <Button className={classes.buttonStyle} color="white">
            v{appVersion}
          </Button>
        </Toolbar>
      </AppBar>
    </div>
  );
}

export default NavBar;
