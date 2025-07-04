@extends('layouts.app')

@section('content')

    @if(session('success_delete'))
        <div class="alert alert-primary alert-dismissible fade show" role="alert">
            {{ session('success_delete') }}
            <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
        </div>
    @endif

    @if(session('failed_delete'))
        <div class="alert alert-warning alert-dismissible fade show" role="alert">
            {{ session('failed_delete') }}
            <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
        </div>
    @endif
    
    <a href="{{ route('customers.create') }}" class="btn btn-primary m-3">Create New Customer</a>
    <h1 class="m-3">Customer Data</h1>
    <table class="table m-3">
        <thead>
            <tr>
                <th scope="col">ID</th>
                <th scope="col">Name</th>
                <th scope="col">Date of Birth</th>
                <th scope="col">Phone Number</th>
                <th scope="col">Email</th>
                <th scope="col">Nationality</th>
                <th scope="col">Action</th>
            </tr>
        </thead>
        <tbody>
            @foreach ($customers as $customer)
                <tr>
                    <th scope="row">{{ $customer['cst_id'] }}</th>
                    <td>{{ $customer['cst_name'] }}</td>
                    <td>{{ $customer['cst_dob'] }}</td>
                    <td>{{ $customer['cst_phone_num'] }}</td>
                    <td>{{ $customer['cst_email'] }}</td>
                    <td>{{ $customer['nationality']['nationality_name'] }}</td>
                    <td>
                        <a href="{{ route('customers.edit', $customer['cst_id']) }}" class="btn btn-warning">Edit</a>
                        <form action="{{ route('customers.destroy', $customer['cst_id']) }}" method="post" onsubmit="return confirm('Are you sure you want to delete this customer?')">
                            @csrf
                            @method('delete')
                            <button type="submit" class="btn btn-danger">Delete</button>
                        </form>
                    </td>
                </tr>
            @endforeach
        </tbody>
    </table>

@endsection
