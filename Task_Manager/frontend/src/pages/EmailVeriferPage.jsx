import { useNavigate } from "react-router-dom"
import {useAuth} from "../context/AuthContext"
import welcomeImage from "../assets/login.png"
import style from "../styles/login.module.css"


const EmailVeriferPage =()=>{
    const {login} = useAuth()
    const navigate = useNavigate()
    
    const handleVerify =()=>{
        //check verification code 

        //navigate to home
       login()
       navigate("/")

    }

    return(
        <div className={style.loginpage}>
            <div className={style.logincard}>
                <h2>Verify Your Email</h2>

                <input 
                    type="password"
                    placeholder="enter verification code" 
                    className={style.inputfield}
                />
                
                <p>Enter the verification code sent to your email to get started</p>

                <button className={style.loginbtn} onClick={handleVerify}>Verify</button>
            </div>
            <img className={style.loginImage} src={welcomeImage} alt="Verify" />
        </div>
    )
}


export default EmailVeriferPage 