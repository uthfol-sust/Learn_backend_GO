import { FiBell, FiUser, FiSearch, FiMenu } from "react-icons/fi"
import style from "../styles/navbar.module.css"
import { useAuth } from "../context/AuthContext"
import { useSidebar } from "../context/SidebarContext"
import { useNavigate } from "react-router-dom";

const Navbar = () => {
    const { isLogin, login} = useAuth();
    const { toggleSidebar,isSidebarOpen  } = useSidebar();
    const navigate = useNavigate(); 

    const handleLogin = ()=>{
        navigate("/login")
        login()
    }

    return (
        <>
        <div className={style.navbar}>
            {!isSidebarOpen && isLogin ?(
                <div className={style.FiMenu}>
                    <FiMenu size={22} onClick={toggleSidebar} style={{ cursor: "pointer" }} />
                    <div>MyTaskMinder</div>
                </div>
            ):(
                <div className={style.FiMenu}>
                    MyTaskMinder
                </div>
            )}

            <div className={style.search}>
                <FiSearch className={style.searchIcon} />
                <input type="search" placeholder="Search tasks..." />
            </div>

            <div className={style.navbarprofile}>
                {isLogin ? (
                    <>
                        <FiBell className={style.icon1} />
                        <FiUser className={style.icon2} />
                    </>
                ) : (
                    <button onClick={handleLogin} className={style.login}>
                        Login
                    </button>
                )}

            </div>
        </div>
        </>
    )
}

export default Navbar