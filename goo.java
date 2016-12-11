Please use this Google doc to code during your interview. To free your hands for coding, we recommend that you use a headset or a phone with speaker option.


Hello! :)

1. input is not necessarily sorted
[1, 2, 5, 7, 10]
[2, 1, 5, 10, 7]
2. int pivot().
    Looking for number 7 in [2, 1, 5, 10, 7]
    pivot() -> [0, 4]
        assume it returned 2, since 7 > 5, now we have [10, 7]
        pivot() -> [3, 4]
        assume it returned 3, since 7 < 10, now we will say 7 can’t be found

        [2, 1, 5, 10, 7] -> [5]
        Looking for 2 inside this array, can 2 always be found? no
        Looking for 1 inside this array, can 1 always be found? no
        ...
        Looking for 7 inside this array, can 7 always be found? no

        input: [2, 1, 5, 10, 7] output: [5]
        input: [1, 2, 5, 7, 10] output: [1, 2, 5, 7, 10]
        input: [10, 7, 5, 2, 1] output: []
        input: null output:

        int[] canBeFoundWithModifiedBinarySearch(int[] input){
            if(input == null || input.length <= 1)
                return input;
            int min = input[input.length - 1];
            int max = input[0];
            ArrayList<Integer> result = ArrayList<Integer>();

            bool[] isValid = bool[input.length];
            for(int i = 0; i < input.length; i++){
                    if(input[i] >= max){
                        isValid[i] = true;
                        max = input[i];
                    }else{
                                isValid[i] = false;
                    }
                }
            for(int i = input.length - 1; i >= 0; i--){
                    if(input[i] <= min){
                                min = input[i];
                    }else{
                        isValid[i] = false;
                    }
                }

            for(int i = 0; i < isValid.length; i++){
                    if(isValid[i]){
                                result.add(input[i])
                    }
                }
            return result.toArray();
        }

