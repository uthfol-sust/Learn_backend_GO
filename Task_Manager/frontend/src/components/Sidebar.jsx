import style from "../styles/sidebar.module.css";
import {useSidebar} from "../context/SidebarContext"
import { TfiAngleDoubleLeft } from "react-icons/tfi";
import { useAuth  } from "../context/AuthContext";

const Sidebar = () => {
  const {closeSidebar , isSidebarOpen} = useSidebar()
  const {logout} = useAuth()

  const handleLogout = () => {
        logout();
    };

  const handleDeshboard=()=>{

  }

  const handleMyTasks=()=>{

  }

  const handleCompleted=()=>{
    
  }

  const handleSetting=()=>{
    
  }


return (
    <div className={`${style.sidebar} ${isSidebarOpen ? style.open : style.close}`}>
        <div className={style.FiMenu}>
            <TfiAngleDoubleLeft size={22} onClick={closeSidebar} style={{ cursor: "pointer" }} />
        </div>
        <h3>Menu</h3>
        <ul>
            <li className={style.menu} onClick={handleDeshboard}>
                Dashboard
            </li>
            <li className={style.menu} onClick={handleMyTasks}>
                My Tasks
            </li>
            <li className={style.menu} onClick={handleCompleted}>
                Completed
            </li>
            <li className={style.menu} onClick={handleSetting}>
                Settings
            </li>
            <li className={style.menu} onClick={handleLogout} >
                Logout
            </li>
        </ul>
    </div>
);
};

export default Sidebar;
