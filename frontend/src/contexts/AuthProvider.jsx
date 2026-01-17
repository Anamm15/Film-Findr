import React, { useMemo } from "react";
import { AuthContext } from "./authContext";
import { useSession } from "../hooks/useUser";

export const AuthProvider = ({ children }) => {
  const {
    data: user,
    isLoading: loading,
    refetch,
  } = useSession();

  const value = useMemo(
    () => ({
      user,
      loading,
      refetchUser: refetch,
    }),
    [user, loading, refetch]
  );

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};
