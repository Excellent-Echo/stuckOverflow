import React, { useState } from "react";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import Typography from "@material-ui/core/Typography";
import InputBase from "@material-ui/core/InputBase";
import MenuItem from "@material-ui/core/MenuItem";
import Menu from "@material-ui/core/Menu";
import MenuIcon from "@material-ui/icons/Menu";
import SearchIcon from "@material-ui/icons/Search";
import AccountCircle from "@material-ui/icons/AccountCircle";
import MoreIcon from "@material-ui/icons/MoreVert";
import { Button, ThemeProvider } from "@material-ui/core";
import Theme from "../styles/Theme";
import userLoginAction from "../redux/user/login/userLoginAction";
import { Redirect, useHistory, useLocation } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import useStyles from "../styles/Style";
import { useEffect } from "react";
import searchQuestionAction from "../redux/question/search/searchQuestionAction";
import SearchBar from "./SearchBar";

const HeaderAppBar = () => {
  const isAuthUser = !!localStorage.getItem("accessToken");
  const history = useHistory();
  const dispatch = useDispatch();
  const classes = useStyles();
  const [anchorEl, setAnchorEl] = React.useState(null);
  const [mobileMoreAnchorEl, setMobileMoreAnchorEl] = React.useState(null);
  const isMenuOpen = Boolean(anchorEl);
  const isMobileMenuOpen = Boolean(mobileMoreAnchorEl);
  const [querySearch, setQuerySearch] = useState("");
  const searchQuestion = useSelector((state) => state.searchQuestion);
  const location = useLocation();

  useEffect(() => {});

  // const handleSearch = (e) => {
  //   setQuerySearch(e.target.value);
  // };

  const handleSubmitSearch = () => {
    dispatch(searchQuestionAction.fetchSearchQuestions(querySearch));
    history.push({
      pathname: "/questions/result/" + querySearch,
    });
  };

  const handleProfileMenuOpen = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMobileMenuClose = () => {
    setMobileMoreAnchorEl(null);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
    handleMobileMenuClose();
  };

  const handleMobileMenuOpen = (event) => {
    setMobileMoreAnchorEl(event.currentTarget);
  };

  const handleLogoutButton = (e) => {
    e.preventDefault();
    dispatch(userLoginAction.logout());
    history.push("/login");
  };

  const menuId = "primary-search-account-menu";
  const renderMenu = (
    <Menu
      anchorEl={anchorEl}
      anchorOrigin={{ vertical: "top", horizontal: "right" }}
      id={menuId}
      keepMounted
      transformOrigin={{ vertical: "top", horizontal: "right" }}
      open={isMenuOpen}
      onClose={handleMenuClose}
    >
      <MenuItem onClick={handleMenuClose}>Profile</MenuItem>
      {isAuthUser && <MenuItem onClick={handleLogoutButton}>Log out</MenuItem>}
    </Menu>
  );

  const mobileMenuId = "primary-search-account-menu-mobile";
  const renderMobileMenu = (
    <Menu
      anchorEl={mobileMoreAnchorEl}
      anchorOrigin={{ vertical: "top", horizontal: "right" }}
      id={mobileMenuId}
      keepMounted
      transformOrigin={{ vertical: "top", horizontal: "right" }}
      open={isMobileMenuOpen}
      onClose={handleMobileMenuClose}
    >
      {!isAuthUser && (
        <MenuItem onClick={() => history.push("/login")}>Log in</MenuItem>
      )}
      {!isAuthUser && (
        <MenuItem onClick={() => history.push("/signup")}>Sign up</MenuItem>
      )}
      {isAuthUser && (
        <MenuItem onClick={handleProfileMenuOpen}>Profile</MenuItem>
      )}
      {isAuthUser && <MenuItem onClick={handleLogoutButton}>Log out</MenuItem>}
    </Menu>
  );

  return (
    <ThemeProvider theme={Theme}>
      <div className={classes.grow}>
        <AppBar position="static">
          <Toolbar>
            {/* <IconButton
              edge="start"
              className={classes.menuButton}
              color="inherit"
              aria-label="open drawer"
            >
              <MenuIcon />
            </IconButton> */}
            <a href="/" style={{ textDecoration: "none", color: "white" }}>
              <Typography
                className={classes.title}
                variant="h6"
                noWrap
                onClick={() => history.push("/")}
              >
                Stuck Overflow
              </Typography>
            </a>
            <SearchBar onChange={(e) => console.log(e.target.value)} />
            {/* <div className={classes.search}>
              <div className={classes.searchIcon}>
                <SearchIcon />
              </div>
              <form onSubmit={handleSubmitSearch}>
                <InputBase
                  placeholder="Search…"
                  classes={{
                    root: classes.inputRoot,
                    input: classes.inputInput,
                  }}
                  inputProps={{ "aria-label": "search" }}
                  onChange={handleSearch}
                />
              </form>
            </div> */}
            <div className={classes.grow} />
            <div className={classes.sectionDesktop}>
              {!isAuthUser && (
                <>
                  <Button
                    href="/login"
                    color="secondary"
                    variant="contained"
                    className={classes.link}
                  >
                    Login
                  </Button>
                  <Button
                    href="/signup"
                    color="secondary"
                    variant="contained"
                    className={classes.link}
                  >
                    Sign up
                  </Button>
                </>
              )}
              {isAuthUser && (
                <IconButton
                  edge="end"
                  aria-label="account of current user"
                  aria-controls={menuId}
                  aria-haspopup="true"
                  onClick={handleProfileMenuOpen}
                  color="inherit"
                >
                  <AccountCircle />
                </IconButton>
              )}
            </div>
            <div className={classes.sectionMobile}>
              <IconButton
                aria-label="show more"
                aria-controls={mobileMenuId}
                aria-haspopup="true"
                onClick={handleMobileMenuOpen}
                color="inherit"
              >
                <MoreIcon />
              </IconButton>
            </div>
          </Toolbar>
        </AppBar>
        {renderMobileMenu}
        {renderMenu}
      </div>
    </ThemeProvider>
  );
};

export default HeaderAppBar;
