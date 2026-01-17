import { Routes, Route } from "react-router-dom";
import LandingPage from "../pages/landing/landing";
import LoginPage from "../pages/login/login";
import RegisterPage from "../pages/register/register";
import DetailFilmPage from "../pages/detail-film/detailFilm";
import DashboardPage from "../pages/dashboard/dashboard";
import ProfilePage from "../pages/profile/Profile";
import ProtectedRoute from "../utils/protectedRoute";
import WatchListPage from "../pages/watchlist/watchlist";
import DashboardLayout from "../layouts/DashboardLayout";
import FilmDashboardPage from "../pages/dashboard/Film";
import TopFilmPage from "../pages/top-film/TopFilm";
import ReviewDashboardPage from "../pages/dashboard/Review";
import GenreDashboardPage from "../pages/dashboard/Genre";
import ContributorDashboardPage from "../pages/dashboard/Contributor";
import AssignmentDashboardPage from "../pages/dashboard/Assignment";
import AddFilmPage from "../pages/dashboard/film/AddFilm";
import UpdateFilmPage from "../pages/dashboard/film/UpdateFilm";

const AppRoutes = () => {
   return (
      <Routes>
         <Route path="/" element={<LandingPage />} />
         <Route path="/login" element={<LoginPage />} />
         <Route path="/register" element={<RegisterPage />} />
         <Route path="/film/:id" element={<DetailFilmPage />} />
         <Route path="/profile/:username" element={<ProfilePage />} />
         <Route path="/weekly-films" element={<TopFilmPage />} />

         <Route
            path="/watchlist"
            element={
               <ProtectedRoute>
                  <WatchListPage />
               </ProtectedRoute>
            }
         />

         <Route
            path="/dashboard"
            element={
               <ProtectedRoute>
                  <DashboardLayout />
               </ProtectedRoute>
            }
         >
            <Route index element={<DashboardPage />} />
            <Route path="films" element={<FilmDashboardPage />} />
            <Route path="reviews" element={<ReviewDashboardPage />} />
            <Route path="genres" element={<GenreDashboardPage />} />
            <Route path="contributors" element={<ContributorDashboardPage />} />
            <Route path="assignments" element={<AssignmentDashboardPage />} />

            <Route path="films/new" element={<AddFilmPage />} />
            <Route path="films/update/:id" element={<UpdateFilmPage />} />
         </Route>
      </Routes>
   );
};

export default AppRoutes;
