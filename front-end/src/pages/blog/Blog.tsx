import Header from "../../components/header/Header";
import styles from "./Blog.module.css";
import "../../Global.css";

function Blog() {
  return (
    <>
      <Header />
      <div className={styles.main}></div>
    </>
  );
}

export default Blog;
