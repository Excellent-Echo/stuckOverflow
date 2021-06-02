import stuckoverflowAPI from "../../../APIs/stuckoverflow";

const resetForm = () => {
  return {
    type: "USER_LOGIN_RESET_FORM",
  };
};

const setUserName = userName => {
  return {
    type: "USER_LOGIN_SET_USERNAME",
    payload: {
      userName: userName,
    },
  };
};

const setPassword = password => {
  return {
    type: "USER_LOGIN_SET_PASSWORD",
    payload: {
      password: password,
    },
  };
};

const setErrorMessage = errorMessage => {
  return {
    type: "USER_LOGIN_SET_ERROR_MESSAGE",
    payload: {
      errorMessage: errorMessage,
    },
  };
};

const setSuccessMessage = successMessage => {
  return {
    type: "USER_LOGIN_SET_SUCCESS_MESSAGE",
    payload: {
      successMessage: successMessage,
    },
  };
};

const startLoading = () => {
  return {
    type: "USER_LOGIN_START_LOADING",
  };
};

const stopLoading = () => {
  return {
    type: "USER_LOGIN_STOP_LOADING",
  };
};

const isLogin = () => {
  return {
    type: "USER_LOGIN_STATUS",
  };
};

const logout = () => {
  return {
    type: "USER_LOGOUT",
  };
};


const login = (userName, password) => async dispatch => {
  try {
    dispatch(startLoading());
    dispatch(setSuccessMessage(""));
    dispatch(setErrorMessage(""));
    const loginData = {
      user_name: userName,
      password: password,
    };

    const user = await stuckoverflowAPI({
      method: "POST",
      url: "/users/login",
      data: loginData,
    });

    const accessToken = user.data.data.token;

    localStorage.setItem("accessToken", accessToken);

    dispatch(isLogin());
    dispatch(stopLoading());

  } catch (error) {
    console.log(error)
    dispatch(setErrorMessage(["The email or password is incorrect."]));
    dispatch(stopLoading());
  }

}

const userLoginAction = {
  resetForm,
  setUserName,
  setPassword,
  login,
  isLogin,
  logout,
};

export default userLoginAction;
