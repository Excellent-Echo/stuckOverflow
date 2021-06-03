import stuckoverflowAPI from "../../../APIs/stuckoverflow";

const resetForm = () => {
  return {
    type: "ASK_QUESTION_RESET_FORM",
  };
};

const setTitle = title => {
  return {
    type: "ASK_QUESTION_SET_TITLE",
    payload: {
      title: title,
    },
  };
};

const setContent = content => {
  return {
    type: "ASK_QUESTION_SET_CONTENT",
    payload: {
      content: content,
    },
  };
};

const setCategory = categoryID => {
  return {
    type: "ASK_QUESTION_SET_CATEGORY",
    payload: {
      categoryID: categoryID,
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

// const setSuccessMessage = successMessage => {
//   return {
//     type: "USER_REGISTER_SET_SUCCESS_MESSAGE",
//     payload: {
//       successMessage: successMessage,
//     },
//   };
// };

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

const askQuestion = (title, content, categoryID) => async dispatch => {
  try {
    dispatch(startLoading());
    // dispatch(setSuccessMessage(""));
    dispatch(setErrorMessage(""));
    const submitData = {
      title: title,
      content: content,
      category_id: categoryID,
    };

    const question = await stuckoverflowAPI({
      method: "POST",
      url: "/questions/ask",
      data: submitData,
      headers: {
        Authorization: localStorage.getItem("accessToken"),
      },
    });

    dispatch(stopLoading());

  } catch (error) {
    console.log(error);
    dispatch(stopLoading());
  }

}

const postQuestionAction = {
  resetForm,
  setTitle,
  setContent,
  setCategory,
  askQuestion,
};

export default postQuestionAction;
