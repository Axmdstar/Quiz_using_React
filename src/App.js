import { useState, useEffect } from "react";
import QuizBrain from './QuizBrain';
import { FaYoutube, FaGithub, FaInstagram } from 'react-icons/fa';

//! Root app 
function App() {

  const [startbtn, setstartbtn] = useState(false);
  const [QuesJson, setQuesJson] = useState(null);
  const [starttext, setstarttext] = useState("START");
  const [title, settitle] = useState("Quiz");
  const [reset, setreset] = useState(false);

  function ResetBtn() {
    setstartbtn(false);
    setQuesJson(null);
    setreset(false);
    setstarttext("START");
  }

  //! Effect Function 
  useEffect(() => {
    if (startbtn) {
      fetch('https://opentdb.com/api.php?amount=10')
        .then((result) => {
          settitle("Quiz")
          return result.json();
        })
        .then((result) => {
          setQuesJson(result.results);
        })
        .catch((error) => {
          setstarttext("Retry")
          settitle("Error")
        })
    }
  }, [startbtn])

  return (
    <div className="mainQuiz">
      <div className="header">

        <h1>{title}</h1>
        <div className="icons">

          <a href="https://www.youtube.com/channel/UCDkUQx2-sfisF8xiiCW_ILw">
            <FaYoutube className="icon" />
          </a>
          <a href="https://instagram.com/devaxmed?igshid=Mzc0YWU1OWY=">
            <FaInstagram className="icon" />
          </a>
          <a href="https://github.com/Axmdstar">
            <FaGithub className="icon" />
          </a>

        </div>


      </div>
      {QuesJson === null
        ?
        <button className="start" onClick={() => { setstartbtn(!startbtn); setstarttext("Loading...") }} > {starttext}</button>
        :
        <div >
          {QuesJson && <QuizBrain Ques={QuesJson} ResetBtn={ResetBtn} />}
        </div>
      }
      {!QuesJson &&  
      <p style={{color:"white", lineHeight:"30px"}}>Introducing my latest project, a simple and fun quiz game! With just a click of a button, my project retrieves 10 random
         questions from a trusted source and presents them to the user. The user can choose to skip questions and come back to them
        later. After answering all the questions, the project displays the final result. It's a great way to test your knowledge
          and have fun at the same time!</p>
          }
    </div>
  );
}

export default App;
