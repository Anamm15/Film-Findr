import React, { useState } from "react";
import { Link, useLocation } from "react-router-dom";
// eslint-disable-next-line no-unused-vars
import { motion, AnimatePresence } from "framer-motion";
import { Menu, X } from "lucide-react";
import { useLogout, useSession } from "../hooks/useUser";
import Button from "./Button";

const Navbar = () => {
   const [isOpen, setIsOpen] = useState(false);
   const location = useLocation();
   const { data: user, isLoading } = useSession()
   const { mutateAsync: logout } = useLogout()

   const navLinks = [
      { name: "Home", href: "/" },
      { name: "Weekly Film", href: "/weekly-film" },
      { name: "Profile", href: "/profile" },
      { name: "Watchlist", href: "/watchlist" },
   ];

   const isActive = (path) => location.pathname === path;

   const handleLogout = async () => {
      try {
         await logout()
         window.location.reload()
      } catch (error) {
         console.log(error)
      }
   }

   return (
      <nav className="fixed top-0 left-0 right-0 z-50 bg-white/80 dark:bg-black/80 backdrop-blur-md border-b border-gray-200 dark:border-white/10">
         <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div className="flex items-center justify-between h-16">
               {/* Logo */}
               <Link
                  to="/"
                  className="font-geist font-bold text-xl tracking-tight bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent"
               >
                  Film-Findr
               </Link>

               {/* Desktop Menu */}
               <div className="hidden md:flex items-center space-x-8">
                  {user && navLinks.map((link) => (
                     <Link
                        key={link.name}
                        to={link.href}
                        className={`relative font-geist font-semibold text-sm transition-colors duration-200 ${isActive(link.href)
                           ? "text-primary dark:text-white"
                           : "text-gray-500 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white"
                           }`}
                     >
                        {link.name}
                        {isActive(link.href) && (
                           <motion.div
                              layoutId="activeNav"
                              className="absolute -bottom-1 left-0 right-0 h-0.5 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full"
                           />
                        )}
                     </Link>
                  ))}

                  {
                     isLoading && user ? (
                        <Button className="text-sm font-semibold px-5" onClick={handleLogout}>Logout</Button>
                     ) : (
                        <div className="flex gap-2">
                           <Link to="/login" className="text-sm font-semibold px-5 border border-gray-200 shadow-lg rounded-full py-2 transition-all duration-200 ease-in-out transform hover:scale-[98%]">Login</Link>
                           <Link to="/register" className="text-sm font-semibold px-5 bg-gradient-to-r from-blue-600 to-purple-600 dark:from-blue-400 dark:to-purple-400 text-white py-2 rounded-full shadow shadow-blue-500 transition-all duration-200 ease-in-out transform hover:scale-[98%]">Register</Link>
                        </div>
                     )
                  }
               </div>

               <div className="md:hidden">
                  <button
                     onClick={() => setIsOpen(!isOpen)}
                     className="text-gray-500 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white p-2"
                  >
                     {isOpen ? <X size={24} /> : <Menu size={24} />}
                  </button>
               </div>
            </div>
         </div>

         <AnimatePresence>
            {isOpen && (
               <motion.div
                  initial={{ opacity: 0, height: 0 }}
                  animate={{ opacity: 1, height: "350px" }}
                  exit={{ opacity: 0, height: 0 }}
                  transition={{ duration: 0.3, ease: "easeInOut" }}
                  className="md:hidden overflow-hidden bg-white dark:bg-black "
               >
                  <div className="flex flex-col items-center justify-center h-full space-y-6">
                     {user && navLinks.map((link) => (
                        <Link
                           key={link.name}
                           to={link.href}
                           onClick={() => setIsOpen(false)}
                           className={`font-geist font-semibold text-lg ${isActive(link.href)
                              ? "text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-purple-600 dark:from-blue-400 dark:to-purple-400"
                              : "text-gray-500 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white"
                              }`}
                        >
                           {link.name}
                        </Link>
                     ))}

                     <div className="w-full flex flex-col gap-4 items-center">
                        {
                           isLoading && user ? (
                              <Button className="font-semibold px-5" onClick={handleLogout}>Logout</Button>
                           ) : (
                              <>
                                 <Link to="/login" className="w-40 font-semibold px-5 border border-gray-200 shadow-lg rounded-md py-2 transition-all duration-200 ease-in-out transform hover:scale-[98%] text-center">Login</Link>
                                 <Link to="/register" className="w-40 font-semibold px-5 bg-gradient-to-r from-blue-600 to-purple-600 dark:from-blue-400 dark:to-purple-400 text-white py-2 rounded-md shadow shadow-blue-500 transition-all duration-200 ease-in-out transform hover:scale-[98%] text-center">Register</Link>
                              </>
                           )
                        }
                     </div>
                  </div>
               </motion.div>
            )
            }
         </AnimatePresence >
      </nav >
   );
};

export default Navbar;