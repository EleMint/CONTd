import React from 'react';
import {
    Redirect,
    Route,
    RouteComponentProps,
    RouteProps
} from 'react-router-dom';

import { useAccount } from '../contexts/Account';

export interface ProtectedRouteProps extends RouteProps {
    fallback?: (props: RouteComponentProps<{}, any, unknown>) => JSX.Element;
    to?: string;
}

export const ProtectedRoute = (props: ProtectedRouteProps) => {
    const { isAuthenticated, isLoading } = useAccount()
    const { component, fallback, to, ...rest } = props;

    const route = React.useMemo(() => {
        if (!isLoading) {
            if (isAuthenticated) {
                return <Route component={component} {...rest} />
            }
            return fallback
                ? <Route component={fallback} {...rest} />
                : <Redirect to={to || "/notfound"} />
        }
        return <>Loading...</>
    }, [component, fallback, isAuthenticated, isLoading, rest, to])

    return route
}
