import { makeStyles } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import NavBar from "./components/NavBar";
import Dashboard from "./pages/Dashboard";
import InitialDataLoader from "./components/InitialDataLoader";

const useStyles = makeStyles((theme) => ({
  root: {
    display: "flex",
    height: "100%",
  },
  content: {
    flex: "1",
    height: "100vh",
    display: "flex",
    paddingTop: theme.spacing(8),
  },
  mainContent: {
    height: "100%",
    width: "100%",
  },
}));

function App() {
  const classes = useStyles();

  return (
    <div className="App">
      <NavBar />
      <main className={classes.content}>
        <div className={classes.mainContent}>
          <InitialDataLoader />
          <Dashboard />
        </div>
      </main>
    </div>
  );
}

export default App;
