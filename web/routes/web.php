<?php

use App\Http\Controllers\CustomerController;
use App\Http\Controllers\FamiliyController;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/
Route::get('/', [CustomerController::class, 'index']);
Route::resource('customers', CustomerController::class);
Route::patch('/families/save/{id}', [FamiliyController::class, 'update'])->name('x.save');
Route::patch('/families/destroy/{id}', [FamiliyController::class, 'destroy'])->name('x.destroy');
