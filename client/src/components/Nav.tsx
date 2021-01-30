import React from 'react';
import { useHistory } from "react-router-dom";
import styled from 'styled-components';

import { useAccount } from '../contexts/Account';

type NavRoute = {
    title: string;
    redirect?: string;
    onClick?: () => void;
}

export const Nav = () => {
    const { isAuthenticated, isLoading, logout } = useAccount()
    const history = useHistory()

    const anonNavRoutes: NavRoute[] = [
        { title: "Login", redirect: "/Login" },
        { title: "Register", redirect: "/Register" }
    ]

    const protectedNavRoutes: NavRoute[] = [
        {
            title: "Logout",
            onClick: () => logout()
        }
    ]

    const mapRoute = (r: NavRoute) => (
        <Route key={r.title}>
            <span
                onClick={() =>
                    r.redirect
                        ? history.push(r.redirect)
                        : r.onClick && r.onClick()
                }
            >{r.title}</span>
        </Route>
    )

    const mapRoutes = (rs: NavRoute[]) => rs.map(r => mapRoute(r))

    return (
        <Base>
            <NavHomeContainer
                onClick={() => history.push("/")}
            >
                CONTd
            </NavHomeContainer>
            <NavRoutesContainer>
                {
                    !isLoading && isAuthenticated
                        ? mapRoutes(protectedNavRoutes)
                        : mapRoutes(anonNavRoutes)
                }
            </NavRoutesContainer>
        </Base>
    )
}

const Base = styled.div`
    position: fixed;
    top: 0;
    left: 0;
    z-index: 1000;
    height: 100px;
    width: 100%;
    background: #008FCC;
`

const NavHomeContainer = styled.div`
    position: absolute;
    height: 100px;
    width: 150px;
    top: 0px;
    left: 0px;
    z-index: 1001;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    background-color: #FFA947;
    color: white;
    cursor: pointer;
`

const NavRoutesContainer = styled.div`
    height: 100px;
    margin: 0 10px;
    float: right;
    list-style-type: none;
    overflow: hidden;
    z-index: 1001;
    display: inline-flex;
`

const Route = styled.div`
    height: 100px;
    margin: 0 10px;
    align-items: center;
    justify-content: center;
    display: inline-flex;
    color: white;
    text-align: center;
    text-decoration: none;
    
    span {
        cursor: pointer;
    }
`
