import { useContext, Fragment } from "react";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import { Link, Typography, Container } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles((theme) => ({
  footer: {
    padding: theme.spacing(3, 2),
    marginTop: "auto",
    backgroundColor: theme.palette.grey[900],
  },
}));

const Footer = observer(({}) => {
  const classes = useStyles();

  const appStore = useContext(AppStoreContext);

  const appVersion = appStore?.appVersion;

  return (
    <Fragment>
      <footer className={classes.footer}>
        <Container maxWidth="sm">
          <center>
            <Typography
              variant="h6"
              style={{ fontFamily: "IBM Plex Mono, mono" }}
            >
              ./onhub-web
            </Typography>
            <br />
            <Typography
              variant="body2"
              color="textSecondary"
              style={{ fontFamily: "IBM Plex Mono, mono" }}
            >
              version {appVersion?.version.slice(0, 8)} <br />
              <Link color="inherit" href="https://github.com/arkits/onhub-web">
                GitHub
              </Link>{"  - "}
              <Link color="inherit" href="/ohw/api/metrics">
                Server Metrics
              </Link>
            </Typography>
          </center>
        </Container>
      </footer>
    </Fragment>
  );
});

export default Footer;
