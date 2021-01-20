import { indigo } from "@material-ui/core/colors";
import { createMuiTheme } from "@material-ui/core/styles";

const theme = createMuiTheme({
  palette: {
    type: "dark",
    primary: {
      main: indigo[500],
    },
    secondary: {
      main: "#f44336",
    },
  },
});

export default theme;
