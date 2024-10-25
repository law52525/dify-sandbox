package integrationtests_test

import (
	"testing"

	"github.com/langgenius/dify-sandbox/internal/core/runner/types"
	"github.com/langgenius/dify-sandbox/internal/service"
)

func TestPythonLargeOutput(t *testing.T) {
	// Test case for base64
	runMultipleTestings(t, 5, func(t *testing.T) {
		resp := service.RunPython3Code(`# declare main function here
def main() -> dict:
    original_strings_with_empty = ["apple", "", "cherry", "date", "", "fig", "grape", "honeydew", "kiwi", "", "mango", "nectarine", "orange", "papaya", "quince", "raspberry", "strawberry", "tangerine", "ugli fruit", "vanilla bean", "watermelon", "xigua", "yellow passionfruit", "zucchini"] * 5

    extended_strings = []

    for s in original_strings_with_empty:
        if s: 
            repeat_times = 600
            extended_s = (s * repeat_times)[:3000]
            extended_strings.append(extended_s)
        else:
            extended_strings.append(s)
    
    return {
        "result": extended_strings,
    }

from json import loads, dumps
from base64 import b64decode

# execute main function, and return the result
# inputs is a dict, and it
inputs = b64decode('e30=').decode('utf-8')
output = main(**loads(inputs))

# convert output to json and print
output = dumps(output, indent=4)

result = f'''<<RESULT>>
{output}
<<RESULT>>'''

print(result)
		`, "", &types.RunnerOptions{
			EnableNetwork: true,
		})
		if resp.Code != 0 {
			t.Fatal(resp)
		}

		if resp.Data.(*service.RunCodeResponse).Stderr != "" {
			t.Fatalf("unexpected error: %s\n", resp.Data.(*service.RunCodeResponse).Stderr)
		}

		if len(resp.Data.(*service.RunCodeResponse).Stdout) != 304487 {
			t.Fatalf("unexpected output, expected 304487 bytes, got %d bytes\n", len(resp.Data.(*service.RunCodeResponse).Stdout))
		}
	})
}


func TestPythonLargeInput(t *testing.T) {
	runMultipleTestings(t, 1000, func(t *testing.T) {
		resp := service.RunPython3Code(`import json
print(json.dumps({
  "srcData": {
    "sex": 1,
    "fat_granule": {
      "category": {
        "milium": {
          "number": 0,
          "level": 1,
          "mask_path": ""
        }
      }
    },
    "orgimage_face_location": [
      486,
      576,
      1357,
      1762,
      2500,
      1875
    ],
    "moisture": {
      "most_serious_class": {
        "face_part": 1,
        "level": 3,
        "class_name": 2
      },
      "overall": {
        "area": 1044,
        "level": 3,
        "area_ratio": 51.97,
        "mask_path": "http://cos.clife.net/clife-ai-cv-skin-application/aQbdeuFE1729766333.png"
      },
      "category": {
        "nose": {
          "level": 3,
          "class_name": 2
        },
        "chin": {
          "level": 2,
          "class_name": 2
        },
        "rightcheek": {
          "level": 3,
          "class_name": 2
        },
        "forehead": {
          "level": 3,
          "class_name": 2
        },
        "leftcheek": {
          "level": 3,
          "class_name": 2
        }
      }
    },
    "basemap_paths": {
      "path_3": "",
      "path_2": "http://cos.clife.net/clife-ai-cv-skin-application/hFdsEAL91729766338.jpg",
      "path_1": "http://cos.clife.net/clife-ai-cv-skin-application/1I6O7FWT1729766337.jpg"
    },
    "skin_type": {
      "class_name": 6
    },
    "face_landmark_list": [
      {
        "x": 559,
        "index": 0,
        "y": 1101
      },
      {
        "x": 556,
        "index": 1,
        "y": 1155
      },
      {
        "x": 553,
        "index": 2,
        "y": 1205
      },
      {
        "x": 549,
        "index": 3,
        "y": 1260
      },
      {
        "x": 553,
        "index": 4,
        "y": 1310
      },
      {
        "x": 556,
        "index": 5,
        "y": 1356
      },
      {
        "x": 563,
        "index": 6,
        "y": 1402
      },
      {
        "x": 570,
        "index": 7,
        "y": 1448
      },
      {
        "x": 583,
        "index": 8,
        "y": 1494
      },
      {
        "x": 600,
        "index": 9,
        "y": 1536
      },
      {
        "x": 621,
        "index": 10,
        "y": 1578
      },
      {
        "x": 648,
        "index": 11,
        "y": 1611
      },
      {
        "x": 675,
        "index": 12,
        "y": 1645
      },
      {
        "x": 706,
        "index": 13,
        "y": 1682
      },
      {
        "x": 744,
        "index": 14,
        "y": 1712
      },
      {
        "x": 791,
        "index": 15,
        "y": 1737
      },
      {
        "x": 846,
        "index": 16,
        "y": 1745
      },
      {
        "x": 917,
        "index": 17,
        "y": 1745
      },
      {
        "x": 975,
        "index": 18,
        "y": 1728
      },
      {
        "x": 1026,
        "index": 19,
        "y": 1703
      },
      {
        "x": 1074,
        "index": 20,
        "y": 1674
      },
      {
        "x": 1122,
        "index": 21,
        "y": 1645
      },
      {
        "x": 1163,
        "index": 22,
        "y": 1611
      },
      {
        "x": 1197,
        "index": 23,
        "y": 1574
      },
      {
        "x": 1221,
        "index": 24,
        "y": 1528
      },
      {
        "x": 1238,
        "index": 25,
        "y": 1482
      },
      {
        "x": 1251,
        "index": 26,
        "y": 1436
      },
      {
        "x": 1265,
        "index": 27,
        "y": 1385
      },
      {
        "x": 1275,
        "index": 28,
        "y": 1331
      },
      {
        "x": 1285,
        "index": 29,
        "y": 1281
      },
      {
        "x": 1292,
        "index": 30,
        "y": 1231
      },
      {
        "x": 1292,
        "index": 31,
        "y": 1176
      },
      {
        "x": 1292,
        "index": 32,
        "y": 1122
      },
      {
        "x": 590,
        "index": 33,
        "y": 1088
      },
      {
        "x": 621,
        "index": 34,
        "y": 1030
      },
      {
        "x": 679,
        "index": 35,
        "y": 1005
      },
      {
        "x": 740,
        "index": 36,
        "y": 1013
      },
      {
        "x": 805,
        "index": 37,
        "y": 1038
      },
      {
        "x": 962,
        "index": 38,
        "y": 1038
      },
      {
        "x": 1033,
        "index": 39,
        "y": 1017
      },
      {
        "x": 1105,
        "index": 40,
        "y": 1013
      },
      {
        "x": 1170,
        "index": 41,
        "y": 1042
      },
      {
        "x": 1214,
        "index": 42,
        "y": 1109
      },
      {
        "x": 880,
        "index": 43,
        "y": 1176
      },
      {
        "x": 873,
        "index": 44,
        "y": 1231
      },
      {
        "x": 866,
        "index": 45,
        "y": 1293
      },
      {
        "x": 863,
        "index": 46,
        "y": 1360
      },
      {
        "x": 778,
        "index": 47,
        "y": 1402
      },
      {
        "x": 825,
        "index": 48,
        "y": 1406
      },
      {
        "x": 859,
        "index": 49,
        "y": 1427
      },
      {
        "x": 897,
        "index": 50,
        "y": 1410
      },
      {
        "x": 951,
        "index": 51,
        "y": 1410
      },
      {
        "x": 645,
        "index": 52,
        "y": 1176
      },
      {
        "x": 686,
        "index": 53,
        "y": 1151
      },
      {
        "x": 757,
        "index": 54,
        "y": 1155
      },
      {
        "x": 791,
        "index": 55,
        "y": 1193
      },
      {
        "x": 750,
        "index": 56,
        "y": 1201
      },
      {
        "x": 682,
        "index": 57,
        "y": 1193
      },
      {
        "x": 975,
        "index": 58,
        "y": 1197
      },
      {
        "x": 1020,
        "index": 59,
        "y": 1164
      },
      {
        "x": 1095,
        "index": 60,
        "y": 1164
      },
      {
        "x": 1135,
        "index": 61,
        "y": 1189
      },
      {
        "x": 1095,
        "index": 62,
        "y": 1205
      },
      {
        "x": 1020,
        "index": 63,
        "y": 1210
      },
      {
        "x": 634,
        "index": 64,
        "y": 1071
      },
      {
        "x": 686,
        "index": 65,
        "y": 1059
      },
      {
        "x": 740,
        "index": 66,
        "y": 1063
      },
      {
        "x": 798,
        "index": 67,
        "y": 1080
      },
      {
        "x": 975,
        "index": 68,
        "y": 1084
      },
      {
        "x": 1033,
        "index": 69,
        "y": 1071
      },
      {
        "x": 1101,
        "index": 70,
        "y": 1067
      },
      {
        "x": 1159,
        "index": 71,
        "y": 1084
      },
      {
        "x": 723,
        "index": 72,
        "y": 1151
      },
      {
        "x": 716,
        "index": 73,
        "y": 1205
      },
      {
        "x": 730,
        "index": 74,
        "y": 1176
      },
      {
        "x": 1057,
        "index": 75,
        "y": 1155
      },
      {
        "x": 1060,
        "index": 76,
        "y": 1214
      },
      {
        "x": 1050,
        "index": 77,
        "y": 1184
      },
      {
        "x": 839,
        "index": 78,
        "y": 1197
      },
      {
        "x": 928,
        "index": 79,
        "y": 1197
      },
      {
        "x": 781,
        "index": 80,
        "y": 1318
      },
      {
        "x": 965,
        "index": 81,
        "y": 1327
      },
      {
        "x": 747,
        "index": 82,
        "y": 1377
      },
      {
        "x": 996,
        "index": 83,
        "y": 1385
      },
      {
        "x": 733,
        "index": 84,
        "y": 1536
      },
      {
        "x": 774,
        "index": 85,
        "y": 1515
      },
      {
        "x": 815,
        "index": 86,
        "y": 1503
      },
      {
        "x": 856,
        "index": 87,
        "y": 1515
      },
      {
        "x": 897,
        "index": 88,
        "y": 1507
      },
      {
        "x": 945,
        "index": 89,
        "y": 1523
      },
      {
        "x": 996,
        "index": 90,
        "y": 1553
      },
      {
        "x": 945,
        "index": 91,
        "y": 1590
      },
      {
        "x": 900,
        "index": 92,
        "y": 1607
      },
      {
        "x": 849,
        "index": 93,
        "y": 1607
      },
      {
        "x": 801,
        "index": 94,
        "y": 1599
      },
      {
        "x": 767,
        "index": 95,
        "y": 1574
      },
      {
        "x": 761,
        "index": 96,
        "y": 1536
      },
      {
        "x": 805,
        "index": 97,
        "y": 1536
      },
      {
        "x": 856,
        "index": 98,
        "y": 1544
      },
      {
        "x": 911,
        "index": 99,
        "y": 1540
      },
      {
        "x": 962,
        "index": 100,
        "y": 1553
      },
      {
        "x": 911,
        "index": 101,
        "y": 1557
      },
      {
        "x": 853,
        "index": 102,
        "y": 1561
      },
      {
        "x": 801,
        "index": 103,
        "y": 1549
      },
      {
        "x": 566,
        "index": 104,
        "y": 1034
      },
      {
        "x": 573,
        "index": 105,
        "y": 984
      },
      {
        "x": 583,
        "index": 106,
        "y": 938
      },
      {
        "x": 600,
        "index": 107,
        "y": 896
      },
      {
        "x": 621,
        "index": 108,
        "y": 858
      },
      {
        "x": 652,
        "index": 109,
        "y": 825
      },
      {
        "x": 686,
        "index": 110,
        "y": 799
      },
      {
        "x": 726,
        "index": 111,
        "y": 783
      },
      {
        "x": 771,
        "index": 112,
        "y": 766
      },
      {
        "x": 815,
        "index": 113,
        "y": 758
      },
      {
        "x": 863,
        "index": 114,
        "y": 753
      },
      {
        "x": 914,
        "index": 115,
        "y": 753
      },
      {
        "x": 965,
        "index": 116,
        "y": 753
      },
      {
        "x": 1020,
        "index": 117,
        "y": 762
      },
      {
        "x": 1071,
        "index": 118,
        "y": 770
      },
      {
        "x": 1118,
        "index": 119,
        "y": 787
      },
      {
        "x": 1159,
        "index": 120,
        "y": 804
      },
      {
        "x": 1193,
        "index": 121,
        "y": 833
      },
      {
        "x": 1224,
        "index": 122,
        "y": 871
      },
      {
        "x": 1248,
        "index": 123,
        "y": 908
      },
      {
        "x": 1265,
        "index": 124,
        "y": 954
      },
      {
        "x": 1279,
        "index": 125,
        "y": 1000
      },
      {
        "x": 1289,
        "index": 126,
        "y": 1051
      }
    ],
    "image_quality": {
      "light_type": 1,
      "blur_type": 1
    },
    "face_shelter": {
      "exist": 0,
      "types": {}
    },
    "pigmentation": {
      "total_pigmentation": {
        "area": 823.0,
        "most_serious_class": 3,
        "level": 2,
        "area_ratio": 0.2578
      },
      "most_serious_class": {
        "area": 347.0,
        "face_part": 4,
        "level": 2,
        "area_ratio": 0.1087,
        "class_name": 3
      },
      "overall": {
        "level": 2,
        "mask_path": "http://cos.clife.net/clife-ai-cv-skin-application/XrwbdZKW1729766334.png"
      },
      "category": {
        "hiding_spot": {
          "area": 55.0,
          "face_part": 1,
          "level": 2,
          "area_ratio": 0.0172
        },
        "freckles": {
          "area": 347.0,
          "face_part": 4,
          "level": 2,
          "area_ratio": 0.1087
        },
        "mole": {
          "area": 155.5,
          "face_part": 3,
          "level": 2,
          "area_ratio": 0.0487
        },
        "chloasma": {
          "area": 265.5,
          "face_part": 3,
          "level": 2,
          "area_ratio": 0.0832
        }
      }
    },
    "oil": {
      "most_serious_class": {
        "face_part": 0,
        "level": 1
      },
      "overall": {
        "area": 0,
        "level": 1,
        "area_ratio": 0.0,
        "mask_path": ""
      },
      "category": {
        "nose": {
          "level": 1
        },
        "chin": {
          "level": 1
        },
        "rightcheek": {
          "level": 1
        },
        "forehead": {
          "level": 1
        },
        "leftcheek": {
          "level": 1
        }
      }
    },
    "acne": {
      "most_serious_class": {
        "number": 4,
        "face_part": 1,
        "level": 2,
        "class_name": 3
      },
      "overall": {
        "number": 4,
        "level": 2,
        "enhence_path": "",
        "mask_path": "http://cos.clife.net/clife-ai-cv-skin-application/DYS5uzV01729766336.png"
      },
      "category": {
        "erythema": {
          "number": 0,
          "face_part": 0,
          "level": 1
        },
        "pustule": {
          "number": 0,
          "face_part": 0,
          "level": 1
        },
        "scar": {
          "number": 0,
          "face_part": 0,
          "level": 1
        },
        "papules": {
          "number": 4,
          "face_part": 1,
          "level": 2
        },
        "nodule": {
          "number": 0,
          "face_part": 0,
          "level": 1
        },
        "pimple": {
          "number": 0,
          "face_part": 0,
          "level": 1
        }
      }
    },
    "face_pose": {
      "yam": 8.96,
      "roll": 3.13,
      "pitch": 5.19
    },
    "isface": 1,
    "facecolor": "A_4",
    "pantone_hex": "#D4A093"
  }
}))`, "", &types.RunnerOptions{
			EnableNetwork: true,
		})
		if resp.Code != 0 {
			t.Fatal(resp)
		}

		if resp.Data.(*service.RunCodeResponse).Stderr != "" {
			t.Fatalf("unexpected error: %s\n", resp.Data.(*service.RunCodeResponse).Stderr)
		}

		if len(resp.Data.(*service.RunCodeResponse).Stdout) == 0 {
			t.Fatalf("unexpected output\n")
		}
	})
}
