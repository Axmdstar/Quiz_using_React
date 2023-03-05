import { useState, useContext, useEffect } from 'react';
import { QuizContext } from './QuizContext';
import he from 'he';


//! Question component
function QusetionBox({ Que }) {

    useEffect(() => {
        const handleexit = (e) => {
            e.preventDefault();
            
            if (document.visibilityState !== "visible") {
                console.log("user left");
                alert("Where did you GO! ");
            }
            
          }
          
          document.addEventListener("visibilitychange", handleexit)
          window.addEventListener('beforeunload', (event) => {
            event.preventDefault();
            event.returnValue = '';
          });
          
    }, []); 


    return (
        <div className='Question' id='question'>
            <p>{ Que && he.decode(Que) }</p>
        </div>
    );
}

//! Bool buttons component 
function Boolbtn({clickfunc, children, boolstate, Answers}) {
    const [isActive, setIsActive] = useState('idle');
    const {Currentindex} = useContext(QuizContext);
    useEffect(() => {
        setIsActive('idle');
    }, [Currentindex]);

    //! click function to check user choice
    function handleClick(e) {
        if (e.target.value === Answers) {
            setIsActive("correct");
            clickfunc("correct");
        }
        else {
            setIsActive("wrong");
            clickfunc("correct");
        }
    }
  
    if (isActive === 'wrong') {
      return (
        <button className='ansbtn wrongbool' value={children} onClick={handleClick} disabled={false}>❌</button>
      );
    } else if (isActive === 'correct') {
      return (
        <button className='ansbtn correctbool' value={children} onClick={handleClick} disabled={false}>&#9989;</button>
      );
    } else {
      return (
        <button className='ansbtn bool' value={children} onClick={handleClick} disabled={boolstate !== 'idle'}>{children}</button>
      );
    }
  }
  
//! main Bool component 
function BooleanBox({ Answers, updataScore }) {
    const [boolstate, setboolstate] = useState('idle');
    const {Currentindex} = useContext(QuizContext);

    useEffect(() => {
        setboolstate('idle');
    }, [Currentindex]);

    const checkAns = (value) => {
        if (value === 'correct') {
            setboolstate('correct');
            updataScore(true, value);
        }
        else {
            setboolstate('wrong');
            updataScore(false,value);
        }
    }
    return (
        <div className='TrueBox'>
            <Boolbtn clickfunc={checkAns} boolstate={boolstate} Answers={Answers}>
                True
            </Boolbtn>
            <Boolbtn clickfunc={checkAns} boolstate={boolstate} Answers={Answers}>
                False
            </Boolbtn>
        </div>
    );
}

//! multiple choice answer buttons, childern component 
function Ansbtn({ children, btnfun, btnstate }) {
    const [onestate, setonestate] = useState('idle');
    let decoded = he.decode(children);

    useEffect(() => {
        if (btnstate === 'idle') {
            setonestate('idle');
        }
    }, [btnstate]);

    //! Function from the main answer component 
    const checkresult = (e) => {
        if (btnfun(e.target.value) === true) {
            setonestate('correct');
        }
        else {
            setonestate('wrong');
        }
      }

    if (onestate === "correct") {
        return (
            <button className='ansbtn correct' value={decoded} onClick={checkresult} disabled={false}  > Correct </button>
        );
    }
    else if (onestate === "wrong") {
        return (
            <button className='ansbtn wrong' value={decoded} onClick={checkresult} disabled={false}  > Wrong </button>
        );
    } 
    else {
        return (
            <>
            {onestate !== btnstate 
                ?(
                    <button className='disable' value={decoded} onClick={checkresult} disabled={true} >{decoded}</button>
                ):(
                    <button className='ansbtn idle' value={decoded} onClick={checkresult} disabled={false} >{decoded}</button>
                )}
            </>
        );
    }
}

//! main multiple component 
function MultipleBox({ Answers, updataScore }) {
    
    const {btnstate, setbtnstate} = useContext(QuizContext);
    const [reset, setreset] = useState(false);
    
    //! function for check user choice 
    const checkAns = (value) => {
        if (value === Answers.corectAns) {
            // add update state btn function
            updataScore(true,value);
            setbtnstate('correct');
            return true;
        }
        else {
            updataScore(false,value);
            setbtnstate('wrong')
            return false;
        }
    }
    
    if (Answers && typeof Answers != "string") {
        return (
            <>
            {"wrong" === btnstate && <p className='correctans'><span style={ {color : 'white' , fontWeight:"lighter"} }>Correct Answer :</span> {Answers.corectAns}</p>}
            <div className='AnswersBox'>
                <Ansbtn btnfun={checkAns} reset={{reset, setreset}} btnstate={btnstate}>{Answers.allAns[0]}</Ansbtn>
                <Ansbtn btnfun={checkAns} reset={{reset, setreset}} btnstate={btnstate}>{Answers.allAns[1]}</Ansbtn>
                <Ansbtn btnfun={checkAns} reset={{reset, setreset}} btnstate={btnstate}>{Answers.allAns[2]}</Ansbtn>
                <Ansbtn btnfun={checkAns} reset={{reset, setreset}} btnstate={btnstate}>{Answers.allAns[3]}</Ansbtn>
            </div>
            </>
        );
    }
}

export {
    BooleanBox,
    MultipleBox,
    QusetionBox
};