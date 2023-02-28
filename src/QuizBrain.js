import { useState, useEffect } from "react";
import he from 'he';
const { MultipleBox, BooleanBox, QusetionBox } = require('./QAs');
const { QuizContext } = require('./QuizContext');

const Results = ({ resultlist, Questions, ResetBtn }) => {
    console.log("Results");
    const emoji = ['🍕','🍔','🍿','🍟','🧀','🥖','🍝','🍧',"🍫","🍾"]
    return (
        <div className="results">
            <table className="resulttable">
            <thead>
                <td></td>
                <td>Question:</td>
                <td>Correct Answer:</td>
                <td>Mark</td>
            </thead>
            {resultlist.map((ele, index) => {
                let decoded = he.decode(Questions[ele.index].question);
                return <tr className="results_p">
                        <td>{emoji[index]}</td>
                        <td >
                            <p key={index}>{decoded}  </p>
                        </td>
                        <td><span>{ele.usrAns}</span></td>
                        <td><span>{ele.checked}</span>  </td>
                    </tr>
            })}

            </table>

            <button className="resetbtn" onClick={ResetBtn}>Play Again</button>
        </div>
    );
}



const QuizBrain = ({ Ques,ResetBtn }) => {
    //! states 
    const [Currentindex, setCurrenntindex] = useState(0);
    const [btnstate, setbtnstate] = useState('idle');
    const [Scorelist, setScorelist] = useState([]);
    const [ScorePp, setScorePp] = useState([]);
    const [Q_A, setQ_A] = useState({ Question: null, Ans: null });


    //! functions 
    function shuffle(Answers) {
        for (let i = Answers.length - 1; i > 0; i--) {
            const j = Math.floor(Math.random() * (i + 1));
            const temp = Answers[i];
            Answers[i] = Answers[j];
            Answers[j] = temp;
        }
        return Answers;
    }

    function updataScore(btnValue, value) {
        if (btnValue) {
            setScorelist(Scorelist => [...Scorelist, "True"]);
            setScorePp(ScorePp => [...ScorePp, {index:Currentindex , usrAns:Ques[Currentindex].correct_answer, checked:"🤩"}])

        }
        else {
            setScorelist(Scorelist => [...Scorelist, "False"]);
            setScorePp(ScorePp => [...ScorePp, {index:Currentindex , usrAns:Ques[Currentindex].correct_answer, checked:"❌"}])

        }
    }

    function Next() {
        if (Currentindex < 9) {
            setCurrenntindex(Currentindex + 1);
            setbtnstate('idle')
        } 
    }

    //! useEffect 
    useEffect(() => {
        if (Ques) {
            setbtnstate('idle')

            if (Ques[Currentindex].type === "boolean") {
                setQ_A({
                    Question: Ques[Currentindex].question,
                    Ans: Ques[Currentindex].correct_answer
                });
            }
            if (Ques[Currentindex].type === "multiple") {
                const arrayOfchooses = [];
                let newarr = arrayOfchooses.concat(Ques[Currentindex].incorrect_answers, Ques[Currentindex].correct_answer);
                setQ_A({
                    Question: Ques[Currentindex].question,
                    Ans: { allAns: shuffle(newarr), corectAns: Ques[Currentindex].correct_answer }
                });
            }
        }
    }, [Currentindex]);

    //! deBug 
    // console.log('Q_A.Question :', Q_A.Question);
    // console.log('Q_A.Ans :', Q_A.Ans);
    // console.log('Score _ List :', Scorelist);

    if (Ques) {
        return (
            <div className="QuizBrain">
                {Currentindex === 9
                    ?(

                        <Results resultlist={ScorePp} Questions={Ques} ResetBtn={ResetBtn} />

                    ):(
                        <QuizContext.Provider value={{ btnstate, setbtnstate, Currentindex }} >
                            <QusetionBox Que={ Q_A.Question } />
                            {Ques[Currentindex].type === "boolean" ?
                                (<BooleanBox updataScore={updataScore} Answers={Q_A.Ans} />) :
                                <MultipleBox Answers={Q_A.Ans} updataScore={updataScore} />}
                            <button type="button" className="Next" onClick={Next}>Next</button>
                        </QuizContext.Provider>

                    )}
            </div>
        );
    }
}


export default QuizBrain;
