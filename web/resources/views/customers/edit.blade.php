@extends('layouts.app')

@section('content')

    <h1 class="mt-5">Customer Detail</h1>
    <form class="mt-5">
        <fieldset disabled>
            <div class="mb-3">
                <label for="cstName" class="form-label">Customer ID</label>
                <input type="text" id="cstName" class="form-control" placeholder={{ $customer['cst_id'] }}>
            </div>
        </fieldset>
        <div class="mb-3">
            <label for="cstName" class="form-label">Name</label>
            <input type="text" class="form-control" id="cstName" value="{{ $customer['cst_name'] }}">
        </div>
        <div class="mb-3">
            <label for="cstDOB" class="form-label">Date of Birth</label>
            <input type="text" class="form-control" id="cstDOB" value="{{ $customer['cst_dob'] }}">
        </div>
        <div class="mb-3">
            <label for="cstPhoneNum" class="form-label">Phone Number</label>
            <input type="text" class="form-control" id="cstPhoneNum" value="{{ $customer['cst_phone_num'] }}">
        </div>
        <div class="mb-3">
            <label for="cstEmail" class="form-label">Email</label>
            <input type="text" class="form-control" id="cstEmail" value="{{ $customer['cst_email'] }}">
        </div>
        <div class="mb-3">
        <label for="nationality_id" class="form-label">Nationality</label>
        <select name="nationality_id" id="nationality_id" class="form-select">
            <option value="">{{ $customer['nationality']['nationality_name'] }}</option>
                @foreach ($nationalities as $nationality)
                    <option value="{{ $nationality['nationality_id'] }}"
                        {{ $nationality['nationality_id'] == ($customer['nationality']['nationality_id'] ?? null) ? 'selected' : '' }}>
                        {{ $nationality['nationality_name'] }}
                    </option>
                @endforeach
        </select>
        </div>
        <button type="submit" class="btn btn-primary">Submit</button>
    </form>

    <div class="container" id="family">
        @foreach ($families as $family)
            {{-- <form action="{{ route('.save') }}" method="POST" class="row">
                @csrf
                <div class="mb-3 col">
                    <label for="fl_name" class="form-label d-flex align-items-center gap-1">
                        Name
                        <span id="{{ $family['fl_id'] . '_edited' }}" style="display:none; color:red">*</span>
                    </label>
                    <input type="text" class="form-control" id="{{ $family['fl_id'] . '_name' }}" value="{{ $family['fl_name'] }}" oninput="editedName({{ $family['fl_id'] }})">
                </div>
                 <div class="mb-3 col">
                    <label for="fl_dob" class="form-label">Date of Birth</label>
                    <input type="date" class="form-control" id="{{ $family['fl_id'] . '_dob' }}" value="{{ $family['fl_Dob'] }}">
                </div>
            </form> --}}
        <form class="row" id="family-form-{{ $family['fl_id'] }}" action="{{ route('x.save', $family['fl_id']) }}" method="POST" onsubmit="return false;">
        @csrf
        @method('PATCH')
            <input type="hidden" name="fl_id" value="{{ $family['fl_id'] }}">
                <div class="mb-3 col">
                    <label for="fl_name" class="form-label d-flex align-items-center gap-1">
                        Name
                        <span id="{{ $family['fl_id'] . '_edited' }}" style="display:none; color:red">*</span>
                    </label>
                    <input type="text" class="form-control" id="{{ $family['fl_id'] . '_name' }}"
                        value="{{ $family['fl_name'] }}" oninput="editedName({{ $family['fl_id'] }})">
                </div>

                <div class="mb-3 col">
                    <label for="fl_dob" class="form-label">Date of Birth</label>
                    <input type="date" class="form-control" id="{{ $family['fl_id'] . '_dob' }}"
                        value="{{ $family['fl_Dob'] }}">
                </div>

                <div class="mb-3 col-auto d-flex align-items-end">
                    <button type="button" onclick="deleteFamily({{ $family['fl_id'] }})" class="btn btn-danger" formaction="{{ route('x.destroy', $family['fl_id']) }}" formmethod="POST">Delete</button> @method('DELETE')
                </div>

                <div class="mb-3 col-auto d-flex align-items-end">
                    <button type="button" onclick="saveFamily({{ $family['fl_id'] }})" class="btn btn-primary">Save</button>
                </div>
        </form>

        @endforeach
    </div>
    <button class="btn btn-primary" onclick="addFamily()">Add New family</button>

    <script>
        function addFamily() {
        const familyContainer = document.getElementById('family');

        const newRow = `
        <form class="row mb-3">
            <div class="col">
                <label for="fl_name" class="form-label">Name</label>
                <input type="text" name="fl_name[]" class="form-control" value="">
            </div>
            <div class="col">
                <label for="fl_dob" class="form-label">Date of Birth</label>
                <input type="date" name="fl_dob[]" class="form-control" value="">
            </div>
        </form>
        `;

            familyContainer.insertAdjacentHTML('beforeend', newRow);
        }

        async function deleteFamily(idFamily) {
            const name = document.getElementById(`${id}_name`).value;
            const dob = document.getElementById(`${id}_dob`).value;
            const form = document.getElementById(`family-form-${id}`);

            const csrfToken = document.querySelector('meta[name="csrf-token"]').content;

            try {
                const response = await fetch(form.action, {
                    method: 'PATCH',
                    headers: {
                        'Content-Type': 'application/json',
                        'X-CSRF-TOKEN': csrfToken,
                        'Accept': 'application/json'
                    },
                    body: JSON.stringify({
                        fl_name: name,
                        fl_dob: dob
                    })
                });

                const data = await response.json();

                if (data.success) {
                    alert('Saved successfully!');
                    document.getElementById(`${id}_edited`).style.display = 'none';
                } else {
                    console.log('Failed to save: ' + (data.message || 'Unknown error') + 'data: ' + (data));
                    console.log(data);
                }
            } catch (error) {
                console.error('Error:', error);
                alert('Something went wrong while saving.');
            }
        }

        function editedName(idFamily) {
            const nameInput = document.getElementById(`${idFamily}_edited`);

            nameInput.style.display = "block";
            console.log(idFamily);
        }

        async function saveFamily(id) {
        const name = document.getElementById(`${id}_name`).value;
        const dob = document.getElementById(`${id}_dob`).value;
        const form = document.getElementById(`family-form-${id}`);

        const csrfToken = document.querySelector('meta[name="csrf-token"]').content;

        try {
            const response = await fetch(form.action, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json',
                    'X-CSRF-TOKEN': csrfToken,
                    'Accept': 'application/json'
                },
                body: JSON.stringify({
                    fl_name: name,
                    fl_dob: dob
                })
            });

            const data = await response.json();

            if (data.success) {
                alert('Saved successfully!');
                document.getElementById(`${id}_edited`).style.display = 'none';
            } else {
                console.log('Failed to save: ' + (data.message || 'Unknown error') + 'data: ' + (data));
                console.log(data);
            }
        } catch (error) {
            console.error('Error:', error);
            alert('Something went wrong while saving.');
        }
    }
        
    </script>
@endsection
