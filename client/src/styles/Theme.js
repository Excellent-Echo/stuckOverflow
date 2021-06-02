import { createMuiTheme } from "@material-ui/core";

const Theme = createMuiTheme({
  palette: {
    primary: {
      light: "#757ce8",
      main: "#3f50b5",
      dark: "#002884",
      contrastText: "#fff",
    },
    secondary: {
      light: "#fffd61",
      main: "#ffca28",
      dark: "#c79a00",
      contrastText: "#000",
    },
  },
});

export default Theme;