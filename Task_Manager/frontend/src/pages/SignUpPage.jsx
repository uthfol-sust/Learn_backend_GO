import { useNavigate } from "react-router-dom"
import signUpImage from "../assets/login.png"
import style from "../styles/login.module.css"


const SignUpPage =()=>{
    const navigate = useNavigate()

    const handleLogin = () =>{
       navigate("/login")
    }

    const handleSignUp =()=>{
        navigate("/emailverify")

    }

    return(
        <div className={style.loginpage}>
            <div className={style.logincard}>
                <h2>Sign Up</h2>
                <p>Register with your valid data to get started</p>

                <label>First Name</label>
                <input 
                    type="text"
                    placeholder="enter your first name"
                    className={style.inputfield}
                />
                <label>Last Name(Optional)</label>
                <input 
                    type="text"
                    placeholder="enter your last name" 
                    className={style.inputfield}
                />

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
                        placeholder="at least 8 characters"
                        className={style.inputfield}
                    />
                </div>

                <label>Confirm your password</label>
                <div className={style.passwordwrapper}>
                    <input 
                        type="password" 
                        placeholder="at least 8 characters"
                        className={style.inputfield}
                    />
                </div>

                <button className={style.loginbtn} onClick={handleSignUp}>Sign Up</button>

                <button className={style.googlebtn}>
                    <img 
                        src="https://upload.wikimedia.org/wikipedia/commons/thumb/c/c1/Google_%22G%22_logo.svg/512px-Google_%22G%22_logo.svg.png"
                        alt="google"
                    />
                    Sign up with Google
                </button>

                <p className={style.registertext}>
                    Already have an account? <span onClick={handleLogin}>Log in</span>
                </p>
            </div>
            <img className={style.loginImage} src={signUpImage} alt="Sign Up" />
        </div>
    )
}


export default SignUpPage