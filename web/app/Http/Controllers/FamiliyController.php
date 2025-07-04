<?php

namespace App\Http\Controllers;

use App\Http\Requests\UpdateFamilyRequest;
use Illuminate\Support\Facades\Http;

class FamiliyController extends Controller {

    /**
     * @param UpdateFamilyRequest $request
     *
     * @return [type]
     */
    public function update(UpdateFamilyRequest $request, int $id)
    {
        $data = $request->validated();

        $save = Http::patch(api_url('families/') . $id, $data);
        if ($save->successful()) {
            return response()->json([
                'success' => true,
                'message' => 'Family updated successfully.',
                'family' => $save->json('data'),
            ]);
        }

        return response()->json([
            'success' => false,
            'message' => $save->json('message'),
            'errors' => $save->json('errors'),
            'er' => $save->body(),
        ], 422);
    }

    public function destroy(int $id)
    {
        $delete = Http::delete(api_url('families/') . $id);
        if ($delete->successful()) {
             return response()->json([
                'success' => true,
                'message' => 'Family deleted successfully.',
                'family' => $delete->json('data'),
            ]);
        }

        return response()->json([
            'success' => false,
            'message' => $delete->json('message'),
            'errors' => $delete->json('errors'),
            'er' => $delete->body(),
        ], 422);
    }
}
