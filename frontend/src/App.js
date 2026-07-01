import {
    BrowserRouter,
    Routes,
    Route,
    Navigate
} from "react-router-dom";

import Home from "./pages/Home";
import Login from "./pages/Login";

function PrivateRoute({children}){

    return sessionStorage.getItem("app_login")==="true"

        ? children

        : <Navigate to="/login"/>

}

function App(){

    return(

        <BrowserRouter>

            <Routes>

                <Route

                    path="/login"

                    element={<Login/>}

                />

                <Route

                    path="/"

                    element={

                        <PrivateRoute>

                            <Home/>

                        </PrivateRoute>

                    }

                />

            </Routes>

        </BrowserRouter>

    );

}

export default App;