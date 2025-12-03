import "./Global.css";
import Header from "./components/header/Header";
import styles from "./App.module.css";
import Footer from "./components/footer/Footer";

function App() {
  return (
    <>
      <Header />
      <div className={styles.main}>
        <div className={styles.textSession1}>
          <h1>Hi, I'm raios!</h1>
          <p>
            This is a blog for my personal toughts and opinions. Inspired by{" "}
            <a className={styles.externalLinks} href="https://lelouch.dev/">
              lelouch.dev
            </a>
            .
          </p>
          <p>
            I'm doing this to pratice my knowledge and my programmer skills, to
            keep it and to share with anyone who wanna it to and to increment my
            resume.
          </p>
          <p>
            Gonna talk about anything I got on my mind and about my progress in
            my career and on "CS major".
          </p>
          <p>
            Want to exchange some ideas?{" "}
            <a className={styles.externalLinks} href="https://x.com/rillminded">
              DM me @rillminded on X
            </a>
          </p>
        </div>
        <div className={styles.textSession2}>
          <h1>Who am I?</h1>
          <p>
            Begginer on Computer Science and IT industry. I study ADS at IFPE
            Campus Garanhuns. Trying to find what really makes me wanna do and
            create things.
          </p>
          <p>I know I do wanna make things.</p>
          <p>
            Love Linux, open-source, catppuccin and doing stuff on terminal.
          </p>
        </div>
      </div>
      <Footer />
    </>
  );
}

export default App;
