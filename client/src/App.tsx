import React from 'react';
import { Route, Switch } from 'react-router-dom';
import styled from 'styled-components';
import { Menu } from './components/Menu';
import { Nav } from './components/Nav';

import { Login } from './routes/Auth/Login';
import { ResetPassword } from './routes/Auth/PasswordReset';
import { Register } from './routes/Auth/Register';
import { RequestResetPassword } from './routes/Auth/RequstPasswordReset';
import { NotFound } from './routes/Error/NotFound';
import { LandingAnon } from './routes/Global/LandingAnon';
import { LandingProtected } from './routes/Global/LandingProctected';
import { ProtectedRoute } from './routes/ProtectedRoute';

export const App = () => {

  return (
    <>
      <Nav />
      <Menu />
      <StyledSwitch>
        <ProtectedRoute exact path={"/"} component={LandingProtected} fallback={LandingAnon} />
        <Route exact path={"/login"} component={Login} />
        <Route exact path={"/register"} component={Register} />
        <Route exact path={"/requestresetpassword"} component={RequestResetPassword} />
        <Route exact path={"/resetpassword"} component={ResetPassword} />

        <Route path={"/notfound"} component={NotFound} />
        <Route path={"/"} component={NotFound} />
      </StyledSwitch>
    </>
  );
}

const StyledSwitch = styled(Switch)`
  position: absolute;
  width: 100%;
  height: 100%;
`
