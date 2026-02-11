import Navbar from "./Navbar";
import { useSidebar } from "../context/SidebarContext";

const Layout = ({ children }) => {
  const { isSidebarOpen } = useSidebar();

  return (
    <div
      style={{
        marginLeft: isSidebarOpen ? "180px" : "0px",
        transition: "margin-left 0.3s ease-in-out"
      }}
    >
      <Navbar />
      <div style={{ padding: "20px" }}>
        {children}
      </div>
    </div>
  );
};

export default Layout;
