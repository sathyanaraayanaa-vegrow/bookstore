import React from 'react'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import LandingPage from '../components/Index';
import SignInSide from '../components/SignIn';
import SignUp from '../components/SignUp';
import BookView from '../components/BookView';
import Borrow from '../components/Borrow';
import CreateBooks from '../components/CreateBooks';
import UpdateBooks from '../components/UpdateBooks';
import Users from '../components/Users';

export default function routes() {
  return (
    <BrowserRouter>
        <Routes>
          <Route path="/" element = {<LandingPage />} />
          <Route path="/signin" element = {<SignInSide />} />
          <Route path="/signup" element = {<SignUp/>} />
          <Route path="/users" element = {<Users />} />
          <Route path="/books" element = {<BookView/>} />
          <Route path="/borrow" element = {<Borrow/>} />
          <Route path="/borrow/:id" element = {<Borrow />} />
          <Route path="/createbook" element = {<CreateBooks />} />
          <Route path="/updatebook/:id" element = {<UpdateBooks />} />
        </Routes>
      </BrowserRouter>
  )
}
