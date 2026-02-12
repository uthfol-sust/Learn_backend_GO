import { BrowserRouter, Routes, Route } from "react-router-dom"
import HomePage from "./pages/HomePage"
import Sidebar from "./components/Sidebar"
import Layout from "./components/Layout";
import LoginPage from "./pages/LoginPage";
import SignUpPage from "./pages/SignUpPage";
import EmailVeriferPage from "./pages/EmailVeriferPage";


function AppContent() {
  return (
    <BrowserRouter>
      <Sidebar />
      <Layout>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/signup" element={<SignUpPage/>} />
          <Route path="/emailverify" element={<EmailVeriferPage/>}/>
        </Routes>
      </Layout>
    </BrowserRouter>
  );
}

function App() {
  return <AppContent />;
}

export default App
