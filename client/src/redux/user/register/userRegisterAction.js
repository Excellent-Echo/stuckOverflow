import stuckoverflowAPI from "../../../APIs/stuckoverflow";

const resetForm = () => {
  return {
    type: "USER_REGISTER_RESET_FORM",
  };
};

const setUserName = userName => {
  return {
    type: "USER_REGISTER_SET_USERNAME",
    payload: {
      userName: userName,
    },
  };
};

const setEmail = email => {
  return {
    type: "USER_REGISTER_SET_EMAIL",
    payload: {
      email: email,
    },
  };
};

const setPassword = password => {
  return {
    type: "USER_REGISTER_SET_PASSWORD",
    payload: {
      password: password,
    },
  };
};

const setErrorMessage = errorMessage => {
  return {
    type: "USER_REGISTER_SET_ERROR_MESSAGE",
    payload: {
      errorMessage: errorMessage,
    },
  };
};

const setSuccessMessage = successMessage => {
  return {
    type: "USER_REGISTER_SET_SUCCESS_MESSAGE",
    payload: {
      successMessage: successMessage,
    },
  };
};

const startLoading = () => {
  return {
    type: "USER_REGISTER_START_LOADING",
  };
};

const stopLoading = () => {
  return {
    type: "USER_REGISTER_STOP_LOADING",
  };
};

const register = (userName, email, password) => async dispatch => {
  try {
    dispatch(startLoading());
    dispatch(setSuccessMessage(""));
    dispatch(setErrorMessage(""));
    const submitData = {
      user_name: userName,
      email: email,
      password: password,
    };

    const user = await stuckoverflowAPI({
      method: "POST",
      url: "/users/register",
      data: submitData,
    });

    dispatch(setSuccessMessage("Congratulations, your account has been successfully created. Please login."))
    dispatch(stopLoading());

  } catch (error) {
    console.log(error);
    dispatch(setErrorMessage(error.response.data.data.error || ["internal server error"]));
    dispatch(stopLoading());
  }

}

const userRegisterAction = {
  resetForm,
  setUserName,
  setEmail,
  setPassword,
  register,
};

export default userRegisterAction;
