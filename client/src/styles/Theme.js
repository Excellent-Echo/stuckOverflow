import { createMuiTheme } from "@material-ui/core";

const Theme = createMuiTheme({
  palette: {
    primary: {
      light: "#484848",
      main: "#212121",
      dark: "#000000",
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