import Header from "../../components/header/Header";
import "../../Global.css";
import styles from "./Login.module.css";

function Login() {
  return (
    <>
      <Header />
      <div className={styles.main}>
        <div className={styles.loginBox}>
          <div className={styles.login}>
            <label htmlFor="login">Username ou Email</label>
            <input type="text" name="login" id="login" />
          </div>
          <div className={styles.password}>
            <label htmlFor="login">Senha</label>
            <input type="password" name="login" id="login" />
          </div>
          <div className={styles.buttons}>
            <a href="/">Entrar</a>
          </div>
        </div>
      </div>
    </>
  );
}

export default Login;
