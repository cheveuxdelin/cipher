import Dashboard from "./pages/Dashboard.jsx";

// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { ref, getDatabase, push, onValue } from 'firebase/database';
import React, { useEffect, useState } from "react";
import { getAuth, createUserWithEmailAndPassword, signInWithEmailAndPassword, onAuthStateChanged, signOut } from "firebase/auth";

import { Navigate, Route, Routes } from "react-router-dom";

import Login from "./pages/Login.jsx";
import Register from "./pages/Register.jsx";

const firebaseConfig = {
  apiKey: "AIzaSyA0xdN4tnUJMo_eDBZYsv33X2Fu7MlGiN4",
  authDomain: "cipher-e7737.firebaseapp.com",
  databaseURL: "https://cipher-e7737-default-rtdb.firebaseio.com",
  projectId: "cipher-e7737",
  storageBucket: "cipher-e7737.appspot.com",
  messagingSenderId: "936238878862",
  appId: "1:936238878862:web:87dad9c8f5ff29cbfc39df"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const database = getDatabase(app);
const auth = getAuth(app);


function signUp(email, password) {

  createUserWithEmailAndPassword(auth, email, password)
    .then((userCredential) => {
      // Signed in 
      const user = userCredential.user;
      console.log(user);
      // ...
    })
    .catch((error) => {
      const errorCode = error.code;
      const errorMessage = error.message;
      console.log(errorCode, errorMessage);
      // ..
    });
}

function login(email, password) {
  signInWithEmailAndPassword(auth, email, password)
    .then((userCredential) => {
      // Signed in 
      const user = userCredential.user;
      console.log(userCredential);
      // ...
    })
    .catch((error) => {
      const errorCode = error.code;
      const errorMessage = error.message;
    });
}

function logout() {
  signOut(auth).then(() => {
    // Sign-out successful.
  }).catch((error) => {
    // An error happened.
  });
}

function cipherSubmitHandler(method, action, text) {
  push(user_ref, {
    created_at: Date.now(),
    method,
    action,
    text,
  });
}



export default function App() {
  const [data, setData] = useState(null);
  const [isUserSignedIn, setIsUserSignedIn] = React.useState(undefined);

  let user_db_ref = ref(database, "/");

  useEffect(() => {

    onAuthStateChanged(auth, user => {
      setIsUserSignedIn(user ? true : false);
    })

    onValue(user_db_ref, (snapshot) => {
      console.log("hola");
      setData(snapshot.val());
    });
  }, []);

  console.log(isUserSignedIn);
  return <React.Fragment>
    {isUserSignedIn != undefined &&
      <Routes>
        <Route path="/dashboard" element={isUserSignedIn ? <Dashboard values={data} cipherSubmitHandler={cipherSubmitHandler} /> : <Navigate to="/login" replace />} />
        <Route path="/login" element={isUserSignedIn ? <Dashboard values={data} cipherSubmitHandler={cipherSubmitHandler} /> : <Login  loginHandler={login}/>} />
        <Route path="/register" element={isUserSignedIn ? <Dashboard values={data} cipherSubmitHandler={cipherSubmitHandler} /> : <Register />} />
        <Route path="*" element={isUserSignedIn ? < Navigate to="/dashboard" replace /> : <Navigate to="/login" replace />} />
      </Routes> 

    }
    < button onClick={logout}>hola</button>

  </React.Fragment>
}

