import {Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import SigninPage from "./pages/patient/SigninPage";
import SignupPage from "./pages/patient/SignupPage";


function App() {
  return (
   
    <Routes>
      <Route path="/" element={<HomePage />}/>
      <Route path="/patient/login" element={<SigninPage/>}/>
      <Route path="/patient/register" element={<SignupPage/>}/>

    </Routes>
  
  );
}

export default App;
