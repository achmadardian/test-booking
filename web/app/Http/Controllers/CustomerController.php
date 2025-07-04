<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class CustomerController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        $fetchCustomers = Http::get(api_url('customers'));

        if ($fetchCustomers->successful()) {
            $customers = $fetchCustomers->json('data');

            return view('customers.index', ['customers' => $customers]);
        }

        return view('customers.index', ['data' => 'this is a customer data']);
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        //
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit(int $id)
    {
        $fetchCustomer = Http::get(api_url('customers/') . $id);

        if (!$fetchCustomer->successful()) {
            return redirect()->route('customers.index');
        }
        $customer = $fetchCustomer->json('data');
        
        $fetchNationalities = Http::get(api_url('nationalities'));
        if (!$fetchCustomer->successful()) {
            return redirect()->route('customers.index');
        }
        $nationalities = $fetchNationalities->json('data');

        $fetchFamilies = Http::get(api_url('customers/') . $id . '/families');
        if (!$fetchFamilies->successful()) {
            return redirect()->route('customers.index');
        }
        $families = $fetchFamilies->json('data');

        foreach ($families as $family) {
            $family['edited'] = false;
        }

        $response = [
            'customer' => $customer,
            'nationalities' => $nationalities,
            'families' => $families,
        ];

        return view('customers.edit', $response);
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy(int $id)
    {
        $deleteCustomer = Http::delete(api_url('customers/'). $id);

        if ($deleteCustomer->successful()) {
            return redirect()->route('customers.index')->with('success_delete', 'Successfully delete customer');
        }

        return redirect()->route('customers.index')->with('failed_delete', 'Failed to delete customer');
    }
}
