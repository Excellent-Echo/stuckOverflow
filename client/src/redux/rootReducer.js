import { combineReducers } from "redux";
import categoryReducer from "./category/get/categoryReducer";
import detailQuestionReducer from "./question/detail/detailQuestionReducer";
import postQuestionReducer from "./question/post/postQuestionReducer";
import questionReducer from "./question/questionReducer";
import searchQuestionReducer from "./question/search/searchQuestionReducer";
import userLoginReducer from "./user/login/userLoginReducer";
import userRegisterReducer from "./user/register/userRegisterReducer";

const rootReducer = combineReducers({
  userRegister: userRegisterReducer,
  userLogin: userLoginReducer,
  question: questionReducer,
  postQuestion: postQuestionReducer,
  detailQuestion: detailQuestionReducer,
  categoryList: categoryReducer,
  searchQuestion: searchQuestionReducer,
});

export default rootReducer;