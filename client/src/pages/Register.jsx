import React, { useContext, useEffect, useState } from "react";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import CssBaseline from "@material-ui/core/CssBaseline";
import TextField from "@material-ui/core/TextField";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Checkbox from "@material-ui/core/Checkbox";
import Link from "@material-ui/core/Link";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import { makeStyles, ThemeProvider } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import Alert from "@material-ui/lab/Alert";
import { useDispatch, useSelector } from "react-redux";
import { useHistory } from "react-router";
import userRegisterAction from "../redux/user/register/userRegisterAction";
import Header from "../components/Header";
import Theme from "../styles/Theme";
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
// function buat styling material ui
const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(6),
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  avatar: {
    margin: theme.spacing(1),
  },
  form: {
    width: "100%",
    marginTop: theme.spacing(3),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const SignUp = () => {
  const userRegisterData = useSelector((state) => state.userRegister);
  const dispatch = useDispatch();
  const history = useHistory();
  const [validated, setValidated] = useState(false);
  const isAuthUser = !!localStorage.getItem("accessToken");
  const classes = useStyles();

  useEffect(() => {
    if (isAuthUser) {
      history.push("/");
    }
    dispatch(userRegisterAction.resetForm());
  }, []);

  const handleRegisterSubmit = (e) => {
    if (isAuthUser) {
      history.push("/");
    }
    e.preventDefault();
    const form = e.currentTarget;
    if (form.checkValidity() === false) {
      e.preventDefault();
      e.stopPropagation();
    } else {
      dispatch(
        userRegisterAction.register(
          userRegisterData.userName,
          userRegisterData.email,
          userRegisterData.password
        )
      );
    }
    setValidated(true);
  };

  return (
    <>
      <ThemeProvider theme={Theme}>
        <Header />
        <Container component="main" maxWidth="xs">
          <CssBaseline />
          <div className={classes.paper}>
            {!userRegisterData.successMessage && (
              <img src={logo} alt="logo" className={classes.avatar} />
            )}
            {!userRegisterData.successMessage && (
              <Typography component="h1" variant="h5">
                Sign up
              </Typography>
            )}
            {userRegisterData.errorMessage && (
              <Alert severity="error">{userRegisterData.errorMessage}</Alert>
            )}
            {!userRegisterData.successMessage && (
              <form
                className={classes.form}
                noValidate
                validated={validated}
                onSubmit={handleRegisterSubmit}
              >
                <Grid container spacing={2}>
                  <Grid item xs={12}>
                    <TextField
                      variant="outlined"
                      required
                      fullWidth
                      id="userName"
                      label="User Name"
                      name="userName"
                      value={userRegisterData.userName}
                      onChange={(e) =>
                        dispatch(userRegisterAction.setUserName(e.target.value))
                      }
                    />
                  </Grid>
                  <Grid item xs={12}>
                    <TextField
                      variant="outlined"
                      required
                      fullWidth
                      id="email"
                      label="Email Address"
                      name="email"
                      autoComplete="email"
                      value={userRegisterData.email}
                      onChange={(e) =>
                        dispatch(userRegisterAction.setEmail(e.target.value))
                      }
                    />
                  </Grid>
                  <Grid item xs={12}>
                    <TextField
                      variant="outlined"
                      required
                      fullWidth
                      name="password"
                      label="Password"
                      type="password"
                      id="password"
                      autoComplete="current-password"
                      value={userRegisterData.password}
                      onChange={(e) =>
                        dispatch(userRegisterAction.setPassword(e.target.value))
                      }
                    />
                  </Grid>
                  <Grid item xs={12}>
                    <FormControlLabel
                      control={
                        <Checkbox
                          value="allowExtraEmails"
                          color="primary"
                          required
                        />
                      }
                      label="Agree to terms and conditions"
                    />
                  </Grid>
                </Grid>
                <Button
                  type="submit"
                  fullWidth
                  variant="contained"
                  color="primary"
                  className={classes.submit}
                  disabled={userRegisterData.isLoading ? true : false}
                >
                  {userRegisterData.isLoading ? "Loading..." : "Sign Up"}
                </Button>
                <Grid container justify="flex-end">
                  <Grid item>
                    <Link href="/login" variant="body2">
                      Already have an account? Log in
                    </Link>
                  </Grid>
                </Grid>
              </form>
            )}
          </div>
          {userRegisterData.successMessage && (
            <Alert severity="success">{userRegisterData.successMessage}</Alert>
          )}
          {!userRegisterData.successMessage && (
            <Box mt={5}>
              <Copyright />
            </Box>
          )}
        </Container>
      </ThemeProvider>
    </>
  );
};

export default SignUp;
