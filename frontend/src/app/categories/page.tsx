"use client";

import { useEffect, useState, useRef } from "react";
import { useRouter } from "next/navigation";
import EditCategoryModal from "@/components/EditCategoryModal";
import Snackbar from "@/components/Snackbar"; // Import the Snackbar component

interface Category {
  Id: string;
  Name: string;
  IsDefault: boolean;
  IsActive: boolean;
}

// Function to generate a consistent background color based on the category name
const getCategoryColor = (categoryName: string) => {
  const colors = [
    "bg-red-500",
    "bg-green-500",
    "bg-blue-500",
    "bg-yellow-500",
    "bg-purple-500",
    "bg-indigo-500",
    "bg-pink-500",
  ];
  let hash = 0;
  for (let i = 0; i < categoryName.length; i++) {
    hash = categoryName.charCodeAt(i) + ((hash << 5) - hash);
  }
  const index = Math.abs(hash) % colors.length;
  return colors[index];
};

export default function CategoriesPage() {
  const [categories, setCategories] = useState<Category[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);
  const router = useRouter();

  // Modal State
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const [newCategoryName, setNewCategoryName] = useState("");
  const [createCategoryError, setCreateCategoryError] = useState<string | null>(null);
  const [isCreatingCategory, setIsCreatingCategory] = useState(false); // Track creation state

  // Edit Modal State
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [editCategoryName, setEditCategoryName] = useState("");
  const [editingCategoryId, setEditingCategoryId] = useState<string | null>(null);
  const [editCategoryError, setEditCategoryError] = useState<string | null>(null);
  const [isEditingCategory, setIsEditingCategory] = useState(false);

  // States to control menus (copied from Dashboard)
  const [menuOpen1, setMenuOpen1] = useState(false);
  const [menuOpen2, setMenuOpen2] = useState(false);
  const menu1Ref = useRef<HTMLUListElement>(null);
  const menu2Ref = useRef<HTMLUListElement>(null);
  const button1Ref = useRef<HTMLButtonElement>(null);
  const button2Ref = useRef<HTMLButtonElement>(null);

  // State for individual category menus
  const [openCategoryMenuId, setOpenCategoryMenuId] = useState<string | null>(
    null,
  );
  const categoryMenuRef = useRef<HTMLDivElement>(null); // Ref for outside click detection

  useEffect(() => {
    const fetchCategories = async () => {
      try {
        const response = await fetch(
          "http://localhost:8080/api/v1/users/categories",
          {
            method: "GET",
            credentials: "include",
          },
        );
        if (!response.ok) {
          if (response.status === 401) {
            router.push("/login");
          }
          const errorData = await response.json();
          throw new Error(errorData.error || "Failed to fetch categories");
        }
        const data = await response.json();
        setCategories(data);
      } catch (err: unknown) {
        if (err instanceof Error) {
          setError(err.message || "Something went wrong");
        }
      }
    };

    fetchCategories();
  }, [router]);

  // useEffect for handling clicks outside the menus (copied from Dashboard)
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        menuOpen1 &&
        menu1Ref.current &&
        !menu1Ref.current.contains(event.target as Node) &&
        button1Ref.current &&
        !button1Ref.current.contains(event.target as Node)
      ) {
        setMenuOpen1(false);
      }
      if (
        menuOpen2 &&
        menu2Ref.current &&
        !menu2Ref.current.contains(event.target as Node) &&
        button2Ref.current &&
        !button2Ref.current.contains(event.target as Node)
      ) {
        setMenuOpen2(false);
      }
      // Handle clicks outside the category menu
      if (
        openCategoryMenuId &&
        categoryMenuRef.current &&
        !categoryMenuRef.current.contains(event.target as Node)
      ) {
        setOpenCategoryMenuId(null);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [menuOpen1, menuOpen2, openCategoryMenuId]);

  // Functions to toggle menus (copied from Dashboard)
  const toggleMenu1 = () => {
    setMenuOpen1(!menuOpen1);
    if (!menuOpen1) {
      setMenuOpen2(false);
    }
  };

  const toggleMenu2 = () => {
    setMenuOpen2(!menuOpen2);
    if (!menuOpen2) {
      setMenuOpen1(false);
    }
  };

  const toggleCategoryMenu = (categoryId: string) => {
    setOpenCategoryMenuId(
      openCategoryMenuId === categoryId ? null : categoryId,
    );
  };

  const handleDeleteCategory = async (category: Category) => {
    if (category.IsDefault) {
      alert("Cannot delete default categories.");
      setOpenCategoryMenuId(null);
      return;
    }

    try {
      const response = await fetch(
        `http://localhost:8080/api/v1/users/categories/${category.Id}`,
        {
          method: "DELETE",
          credentials: "include",
        },
      );

      if (!response.ok) {
        const errorData = await response.json();
        setError(errorData.error || "Failed to delete category");
        return;
      }
      setError(null);
      setSuccess(`Category "${category.Name}" deleted successfully.`);
      // Update the categories list after successful deletion
      setCategories(categories.map(c => c.Id === category.Id ? {...c, IsActive: false} : c));
      setOpenCategoryMenuId(null);
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message || "Something went wrong during deletion");
      }
    }
  };

  const handleReactivateCategory = async (category: Category) => {
    try {
      const response = await fetch(
        `http://localhost:8080/api/v1/users/categories/${category.Id}/reactivate`,
        {
          method: "POST",
          credentials: "include",
        },
      );

      if (!response.ok) {
        const errorData = await response.json();
        setError(errorData.error || "Failed to reactivate category");
        return;
      }
      setError(null);
      setSuccess(`Category "${category.Name}" reactivated successfully.`);
      // Update the categories list after successful reactivation
      setCategories(categories.map(c => c.Id === category.Id ? {...c, IsActive: true} : c));
      setOpenCategoryMenuId(null);
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message || "Something went wrong during reactivation");
      }
    }
  };

  const openAddModal = () => {
    setIsAddModalOpen(true);
  };

  const closeAddModal = () => {
    setIsAddModalOpen(false);
    setNewCategoryName("");
    setCreateCategoryError(null);
    setIsCreatingCategory(false); // Reset creating state
  };

  const handleCreateCategory = async () => {
    setCreateCategoryError(null);
    setIsCreatingCategory(true); // Disable the button
    try {
      const response = await fetch(
        "http://localhost:8080/api/v1/users/categories",
        {
          method: "POST",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ name: newCategoryName }),
        },
      );

      if (!response.ok) {
        const errorData = await response.json();
        setIsCreatingCategory(false); // Re-enable the button
        throw new Error(errorData.error || "Failed to create category");
      }

      const newCategory = await response.json();
      setCategories([...categories, newCategory]);
      setSuccess(`Category "${newCategory.Name}" created successfully.`);
      closeAddModal();
    } catch (err: unknown) {
      setIsCreatingCategory(false); // Ensure button is re-enabled even on error
      if (err instanceof Error) {
        setCreateCategoryError(err.message);
      }
    }
  };

  const openEditModal = (category: Category) => {
    setEditingCategoryId(category.Id);
    setEditCategoryName(category.Name);
    setIsEditModalOpen(true);
  };

  const closeEditModal = () => {
    setIsEditModalOpen(false);
    setEditingCategoryId(null);
    setEditCategoryName("");
    setEditCategoryError(null);
    setIsEditingCategory(false);
  };

  const handleEditCategory = async (id: string, newName: string) => {
    setEditCategoryError(null);
    setIsEditingCategory(true);
    try {
      const response = await fetch(
        `http://localhost:8080/api/v1/users/categories/${id}`,
        {
          method: "PATCH",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ name: newName }),
        },
      );

      if (!response.ok) {
        const errorData = await response.json();
        setIsEditingCategory(false);
        throw new Error(errorData.error || "Failed to edit category");
      }

      const updatedCategory = await response.json();
      setCategories(categories.map(cat => cat.Id === updatedCategory.Id ? updatedCategory : cat));
      setSuccess(`Category "${updatedCategory.Name}" updated successfully.`);
      closeEditModal();
    } catch (err: unknown) {
      setIsEditingCategory(false);
      if (err instanceof Error) {
        setEditCategoryError(err.message);
      }
    }
  };

  const handleCloseSnackbar = () => {
    setError(null);
    setSuccess(null);
    setCreateCategoryError(null);
    setEditCategoryError(null);
  };

  useEffect(() => {
    if (success || error || createCategoryError || editCategoryError) {
      const timer = setTimeout(() => {
        handleCloseSnackbar();
      }, 3000); // Adjust the duration as needed
      return () => clearTimeout(timer);
    }
  }, [success, error, createCategoryError, editCategoryError]);

  const activeCategories = categories.filter((cat) => cat.IsActive);
  const inactiveCategories = categories.filter((cat) => !cat.IsActive);

  return (
    <div className="h-screen bg-gray-900 p-6 text-gray-100">
      {/* Render snackbar if there's an error or success message */}
      {error && <Snackbar type="error" message={error} onClose={handleCloseSnackbar} />}
      {success && <Snackbar type="success" message={success} onClose={handleCloseSnackbar} />}
      {createCategoryError && <Snackbar type="error" message={createCategoryError} onClose={handleCloseSnackbar} />}
      {editCategoryError && <Snackbar type="error" message={editCategoryError} onClose={handleCloseSnackbar} />}

      {/* Navbar (copied from Dashboard) */}
      <div className="navbar bg-gray-800 text-gray-100 shadow-lg rounded-lg mb-6 relative">
        <div className="flex-none relative">
          <button
            ref={button1Ref}
            className="btn btn-square btn-ghost"
            onClick={toggleMenu1}
          >
            {/* First menu toggle button */}
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              className="inline-block h-5 w-5 stroke-current"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M4 6h16M4 12h16M4 18h16"
              ></path>
            </svg>
          </button>

          {/* Conditionally render your first menu */}
          {menuOpen1 && (
            <ul
              ref={menu1Ref}
              className="menu bg-gray-700 rounded-box w-56 absolute top-12 left-0 z-50 shadow"
            >
              <li>
                <a
                  onClick={() => router.push("/dashboard")}
                  className="hover:bg-gray-600 cursor-pointer"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-5 w-5 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
                    />
                  </svg>
                  Transactions
                </a>
              </li>
              <li>
                <a className="hover:bg-gray-600">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-5 w-5 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                  </svg>
                  Categories
                </a>
              </li>
              <li>
                <a
                  onClick={() => router.push("/insights")}
                  className="hover:bg-gray-600 cursor-pointer"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-5 w-5 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
                    />
                  </svg>
                  Insights
                </a>
              </li>
            </ul>
          )}
        </div>
        <div className="flex-1">
          <a className="btn btn-ghost text-xl font-mono">tracker</a>
        </div>

        <div className="flex-none relative">
          <button
            ref={button2Ref}
            className="btn btn-square btn-ghost"
            onClick={toggleMenu2}
          >
            {/* Second menu toggle button */}
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              className="inline-block h-5 w-5 stroke-current"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M5 12h.01M12 12h.01M19 12h.01M6 12a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0z"
              ></path>
            </svg>
          </button>

          {/* Conditionally render your second menu */}
          {menuOpen2 && (
            <ul
              ref={menu2Ref}
              className="menu bg-gray-700 rounded-box w-56 absolute top-12 right-0 z-50 shadow"
            >
              <li>
                <a className="hover:bg-gray-600">Profile</a>
              </li>
              <li>
                <a className="hover:bg-gray-600">Settings</a>
              </li>
              <li>
                <a className="hover:bg-gray-600">Sign out</a>
              </li>
            </ul>
          )}
        </div>
      </div>

      <div className="max-w-4xl mx-auto bg-gray-800 rounded-lg shadow-md p-8">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-2xl font-bold font-mono">Manage Categories</h2>
          <button
            onClick={openAddModal}
            className="bg-gray-700 hover:bg-gray-600 text-gray-100 font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline font-mono"
          >
            Add Category
          </button>
        </div>

        {categories.length === 0 && !error && (
          <p className="text-gray-400 font-mono">Loading categories...</p>
        )}

        {categories.length > 0 && (
          <div>
            <h3 className="text-lg font-bold mb-2 font-mono">Active Categories</h3>
            <div className="space-y-4 mb-6">
              {activeCategories.map((category) => (
                <div
                  key={category.Id}
                  className="flex items-center justify-between p-4 border border-gray-700 rounded-lg shadow-sm hover:shadow-lg transition-shadow bg-gray-800 hover:bg-gray-700"
                >
                  <div className="flex items-center justify-center space-x-4">
                    <div
                      className={`w-10 h-10 flex items-center justify-center rounded-full font-bold font-mono uppercase text-gray-100 ${getCategoryColor(category.Name)}`}
                    >
                      {category.Name.charAt(0)}
                    </div>
                    <p className="font-mono">{category.Name}</p>
                  </div>
                  <div className="relative">
                    <button
                      onClick={() => toggleCategoryMenu(category.Id)}
                      className="btn btn-ghost btn-sm"
                      disabled={category.IsDefault} // Disable menu for default categories
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        className="inline-block w-5 h-5 stroke-current"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth="2"
                          d="M5 12h.01M12 12h.01M19 12h.01M6 12a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0z"
                        ></path>
                      </svg>
                    </button>
                    {openCategoryMenuId === category.Id && (
                      <div
                        ref={categoryMenuRef}
                        className="absolute right-0 mt-2 w-36 bg-gray-700 rounded-md shadow-lg z-10"
                      >
                        {!category.IsDefault && (
                          <a
                            href="#"
                            className="block px-4 py-2 text-sm text-gray-100 hover:bg-gray-600 font-mono"
                            onClick={() => {
                              openEditModal(category);
                              setOpenCategoryMenuId(null);
                            }}
                          >
                            Edit
                          </a>
                        )}
                        {!category.IsDefault && ( // Conditionally render delete
                          <a
                            href="#"
                            className="block px-4 py-2 text-sm text-gray-100 hover:bg-gray-600 font-mono"
                            onClick={() => handleDeleteCategory(category)}
                          >
                            Delete
                          </a>
                        )}
                        {category.IsDefault && (
                          <span className="block px-4 py-2 text-sm text-gray-400 cursor-default font-mono">
                            Default
                          </span>
                        )}
                      </div>
                    )}
                  </div>
                </div>
              ))}
            </div>

            <h3 className="text-lg font-bold mb-2 font-mono">Inactive Categories</h3>
            <div className="space-y-4">
              {inactiveCategories.map((category) => (
                <div
                  key={category.Id}
                  className="flex items-center justify-between p-4 border border-gray-700 rounded-lg shadow-sm hover:shadow-lg transition-shadow bg-gray-800 hover:bg-gray-700 opacity-50 saturate-0 grayscale"
                >
                  <div className="flex items-center justify-center space-x-4">
                    <div
                      className={`w-10 h-10 flex items-center justify-center rounded-full font-bold font-mono uppercase text-gray-100 ${getCategoryColor(category.Name)}`}
                    >
                      {category.Name.charAt(0)}
                    </div>
                    <p className="font-mono">{category.Name}</p>
                  </div>
                  <div className="relative">
                    <button
                      onClick={() => toggleCategoryMenu(category.Id)}
                      className="btn btn-ghost btn-sm"
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        className="inline-block w-5 h-5 stroke-current"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth="2"
                          d="M5 12h.01M12 12h.01M19 12h.01M6 12a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0z"
                        ></path>
                      </svg>
                    </button>
                    {openCategoryMenuId === category.Id && (
                      <div
                        ref={categoryMenuRef}
                        className="absolute right-0 mt-2 w-36 bg-gray-700 rounded-md shadow-lg z-10"
                      >
                        {!category.IsDefault && (
                          <a
                            href="#"
                            className="block px-4 py-2 text-sm text-gray-100 hover:bg-gray-600 font-mono"
                            onClick={() => {
                              openEditModal(category);
                              setOpenCategoryMenuId(null);
                            }}
                          >
                            Edit
                          </a>
                        )}
                        <a
                          href="#"
                          className="block px-4 py-2 text-sm text-gray-100 hover:bg-gray-600 font-mono"
                          onClick={() => handleReactivateCategory(category)}
                        >
                          Reactivate
                        </a>
                      </div>
                    )}
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>

      {/* Add Category Modal */}
      <dialog id="add_category_modal" className="modal modal-bottom sm:modal-middle backdrop:bg-gray-900 backdrop:bg-opacity-50" open={isAddModalOpen}>
        <div className="modal-box bg-gray-800 text-gray-100">
          <h3 className="font-bold text-lg font-mono">Add New Category</h3>
          <form method="dialog" className="mt-4">
            {/* if there is a button in form, it will close the modal */}
            <div className="mb-4">
              <label htmlFor="categoryName" className="block text-gray-300 text-sm font-bold mb-2 font-mono">
                Category Name
              </label>
              <input
                type="text"
                id="categoryName"
                placeholder="Enter category name"
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-100 leading-tight focus:outline-none focus:shadow-outline bg-gray-700 border-gray-600 font-mono"
                value={newCategoryName}
                onChange={(e) => setNewCategoryName(e.target.value)}
              />
            </div>
            {createCategoryError && <Snackbar type="error" message={createCategoryError} onClose={handleCloseSnackbar} />}
            <div className="modal-action">
              <button className="btn btn-error font-mono" onClick={closeAddModal}>Cancel</button>
              <button
                className="btn btn-primary font-mono"
                onClick={handleCreateCategory}
                disabled={!newCategoryName.trim() || isCreatingCategory} // Disable during creation
              >
                Create
              </button>
            </div>
          </form>
        </div>
      </dialog>

      {/* Edit Category Modal */}
      {editingCategoryId && (
        <EditCategoryModal
          isOpen={isEditModalOpen}
          onClose={closeEditModal}
          onEdit={handleEditCategory}
          editCategoryError={editCategoryError}
          isEditingCategory={isEditingCategory}
          handleCloseSnackbar={handleCloseSnackbar}
          categoryName={editCategoryName}
          categoryId={editingCategoryId}
        />
      )}
    </div>
  );
}
