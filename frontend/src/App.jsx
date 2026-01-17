import { AuthProvider } from "./contexts/authProvider";
import AppRoutes from "./routes/route";
import Navbar from "./components/navbar";

function App() {
  return (
    <AuthProvider >
      <Navbar />
      <AppRoutes />
    </AuthProvider>
  );
}

export default App;
