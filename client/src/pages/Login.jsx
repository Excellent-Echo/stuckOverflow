import React, { useContext, useEffect } from "react";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import CssBaseline from "@material-ui/core/CssBaseline";
import TextField from "@material-ui/core/TextField";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Checkbox from "@material-ui/core/Checkbox";
import Link from "@material-ui/core/Link";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";
import Typography from "@material-ui/core/Typography";
import { makeStyles, ThemeProvider } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import Alert from "@material-ui/lab/Alert";
import { useDispatch, useSelector } from "react-redux";
import { useHistory } from "react-router";
import userLoginAction from "../redux/user/login/userLoginAction";
import Header from "../components/Header";
import Theme from "../styles/Theme";
import { LoadingContext } from "react-router-loading";
import logo from "../assets/stuckoverflowlogo.png";

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {"Copyright © "}
      <Link color="inherit" href="https://material-ui.com/">
        Stuck Overflow
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
}

// function styling material ui
const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(),
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  avatar: {
    margin: theme.spacing(1),
  },
  form: {
    width: "100%", // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const SignIn = () => {
  const userLoginData = useSelector((state) => state.userLogin);
  const dispatch = useDispatch();
  const history = useHistory();
  const isAuthUser = !!localStorage.getItem("accessToken");
  const classes = useStyles();
  const loadingContext = useContext(LoadingContext);

  useEffect(() => {
    if (isAuthUser) {
      history.push("/");
    }
    dispatch(userLoginAction.resetForm());
    loadingContext.done();
  }, []);

  const handlerLoginSubmit = (e) => {
    e.preventDefault();
    dispatch(
      userLoginAction.login(userLoginData.userName, userLoginData.password)
    );
    history.push("/");
  };

  return (
    <>
      <ThemeProvider theme={Theme}>
        <Header />
        <Container component="main" maxWidth="xs">
          <CssBaseline />
          <div className={classes.paper}>
            <img src={logo} alt="logo" className={classes.avatar} />
            <Typography component="h1" variant="h5">
              Log in
            </Typography>
            {userLoginData.errorMessage && (
              <Alert severity="error">{userLoginData.errorMessage}</Alert>
            )}
            {!userLoginData.successMessage && (
              <form
                className={classes.form}
                noValidate
                onSubmit={handlerLoginSubmit}
              >
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  id="userName"
                  label="User Name"
                  name="userName"
                  value={userLoginData.userName}
                  onChange={(e) =>
                    dispatch(userLoginAction.setUserName(e.target.value))
                  }
                />
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  name="password"
                  label="Password"
                  type="password"
                  id="password"
                  value={userLoginData.password}
                  onChange={(e) =>
                    dispatch(userLoginAction.setPassword(e.target.value))
                  }
                />
                <FormControlLabel
                  control={<Checkbox value="remember" color="primary" />}
                  label="Remember me"
                />
                <Button
                  type="submit"
                  fullWidth
                  variant="contained"
                  color="primary"
                  className={classes.submit}
                  disabled={userLoginData.isLoading ? true : false}
                >
                  {userLoginData.isLoading ? "Loading..." : "Login"}
                </Button>
                <Grid container>
                  <Grid item xs>
                    <Link href="#" variant="body2">
                      Forgot password?
                    </Link>
                  </Grid>
                  <Grid item>
                    <Link href="#" variant="body2">
                      {"Don't have an account? Sign Up"}
                    </Link>
                  </Grid>
                </Grid>
              </form>
            )}
          </div>
          {userLoginData.successMessage && (
            <Alert severity="success">{userLoginData.successMessage}</Alert>
          )}
          <Box mt={8}>
            <Copyright />
          </Box>
        </Container>
      </ThemeProvider>
    </>
  );
};

export default SignIn;
