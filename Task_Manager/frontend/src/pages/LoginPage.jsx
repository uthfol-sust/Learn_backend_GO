import { useNavigate } from "react-router-dom"
import loginImage from "../assets/login.png"
import style from "../styles/login.module.css"


const LoginPage =()=>{
    const navigate = useNavigate()
    
    const handleRegister = ()=>{
        navigate("/signup")
    }

    return(
        <div  className={style.loginpage}>
              <div className={style.logincard}>
                    <h2>Log in.</h2>
                    <p>Log in with your data that you entered during your registration</p>

                    <label>Enter your email address</label>
                    <input 
                    type="email" 
                    placeholder="name@example.com"
                    className={style.inputfield}
                    />

                    <label>Enter your password</label>
                    <div className={style.passwordwrapper}>
                    <input 
                        type="password" 
                        placeholder="atleast 8 characters"
                        className={style.inputfield}
                    />
                    </div>

                    <p className={style.forgotpassword}>Forgot password?</p>

                    <button className={style.loginbtn}>Log in</button>

                    <button className={style.googlebtn}>
                    <img 
                        src="https://upload.wikimedia.org/wikipedia/commons/thumb/c/c1/Google_%22G%22_logo.svg/512px-Google_%22G%22_logo.svg.png"
                        alt="google"
                    />
                    Sign in with Google
                    </button>

                    <p className={style.registertext}>
                    Don't have an account? <span onClick={handleRegister}>Register</span>
                    </p>
                </div>
             <img className={style.loginImage} src={loginImage} alt="Login" />
        </div>
    )
}

export default LoginPage