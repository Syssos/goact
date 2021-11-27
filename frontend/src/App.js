import './App.css';
import Chatterbox from './components/Chatterbox';
import LoginForm from './components/LoginForm';

function getCookie(name) {
    var dc = document.cookie;
    var prefix = name + "=";
    var begin = dc.indexOf("; " + prefix);
    if (begin === -1) {
        begin = dc.indexOf(prefix);
        if (begin !== 0) return null;
    }
    else
    {
        begin += 2;
        var end = document.cookie.indexOf(";", begin);
        if (end === -1) {
        end = dc.length;
        }
    }
    // because unescape has been deprecated, replaced with decodeURI
    //return unescape(dc.substring(begin + prefix.length, end));
    return decodeURI(dc.substring(begin + prefix.length, end));
}

function App() {

  if(!getCookie("token")) return (<div className="App"><LoginForm /></div>);
  
  return (
    <div className="App">
      <Chatterbox />
    </div>
  );
}

export default App;
