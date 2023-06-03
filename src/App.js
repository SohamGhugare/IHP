import {Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import SigninPage from "./pages/patient/SigninPage";
import SignupPage from "./pages/patient/SignupPage";
import SigninPageDoc from "./pages/doctor/SigninPage";
import SignupPageDoc from "./pages/doctor/SignupPage";
import Profile from "./pages/doctor/Profile";


function App() {
  return (
   
    <Routes>
      <Route path="/" element={<HomePage />}/>
      <Route path="/patient/login" element={<SigninPage/>}/>
      <Route path="/patient/register" element={<SignupPage/>}/>
      <Route path="/doctor/login" element={<SigninPageDoc/>}/>
      <Route path="/doctor/register" element={<SignupPageDoc/>}/>
      <Route path="/doctor/profile" element={<Profile/>}/>

    </Routes>
  
  );
}

export default App;
